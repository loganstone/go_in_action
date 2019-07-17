package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var (
	// ErrTimeout is returned when a value is
	// received on the timeout channel.
	ErrTimeout = errors.New("received timeout")
	// ErrInterrupt is returned when an event
	// from the OS is received.
	ErrInterrupt = errors.New("received interrupt")
)

// Runner runs a set of tasks within a given timeout and can be
// shut down on an operating system interrupt.
type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	tasks     []func(int)
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.goInterrupt() {
			return ErrInterrupt
		}
		// It does not stop with "os.Interrupt",
		// while processing task.
		// When an interrupt event is sent,
		// program terminate before next task.
		task(id)
	}
	return nil
}

func (r *Runner) goInterrupt() bool {
	select {
	// Signaled when an interrupt event is sent.
	case <-r.interrupt:
		// Stop receiving any further signals.
		signal.Stop(r.interrupt)
		return true
	// Continue running as normal.
	default:
		return false
	}
}

// Add attaches tasks to the Runner.
// A task is a function that takes an int ID.
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start runs all tasks and monitors channel events.
func (r *Runner) Start() error {
	// We want to receive all interrupt based signals.
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	// Wait until processing is done or timeout.
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

// New returns a new ready-to-use Runner.
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}
