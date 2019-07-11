package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 11
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad)
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("작업 : %d", post)
	}

	close(tasks)

	wg.Wait()
}

func worker(tasks chan string, n int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("작업자: %d: 종료합니다.\n", n)
			return
		}

		fmt.Printf("작업자: %d: 작업 시작: %s.\n", n, task)
		do := func() {
			sleep := rand.Int63n(100)
			time.Sleep(time.Duration(sleep) * time.Millisecond)
		}
		do()
		fmt.Printf("작업자: %d: 작업 완료: %s.\n", n, task)
	}
}
