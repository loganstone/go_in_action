package work

import "sync"

// Worker .
type Worker interface {
	Task()
}

// Pool .
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New .
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			// It is blocked until channel is closed.
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

// Run .
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// ShutDown .
func (p *Pool) ShutDown() {
	close(p.work)
	p.wg.Wait()
}
