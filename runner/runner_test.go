package runner

import (
	"fmt"
	"log"
	"time"
)

const timeout = 3 * time.Second

func ExampleRunner() {

	createTask := func() func(int) {
		return func(id int) {
			fmt.Printf("Processor - Task #%d.\n", id)
			time.Sleep(time.Duration(id) * time.Second)
		}
	}

	log.Println("Starting work.")

	r := New(timeout)
	r.Add(createTask(), createTask(), createTask())

	// Run the tasks and handle the result.
	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeout:
			fmt.Println("Terminating due to timeout.")
		case ErrInterrupt:
			fmt.Println("Terminating due to interrupt.")
		}
	}

	log.Println("Process ended.")
	// Output:
	// Processor - Task #0.
	// Processor - Task #1.
	// Processor - Task #2.
	// Terminating due to timeout.
}
