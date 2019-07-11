package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

const maxRunners = 4

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	baton := make(chan int)
	wg.Add(1)

	go Runner(baton)

	baton <- 1

	wg.Wait()
}

// Runner .
func Runner(baton chan int) {
	var newRunner int
	runner := <-baton
	if runner == 1 {
		fmt.Printf("탕! %d 번째 주자가 달리기 시작했습니다.\n", runner)
	} else {
		fmt.Printf("%d 번째 주자가 바통을 받아 달리기 시작했습니다.\n", runner)
	}

	if runner != maxRunners {
		newRunner = runner + 1
		fmt.Printf("%d 번째 주자가 대기 합니다.\n", newRunner)
		go Runner(baton)
	}

	running := func() {
		time.Sleep(100 * time.Millisecond)
	}
	running()

	if runner == maxRunners {
		fmt.Printf("%d 번째 주자가 도착 했습니다. 경기가 끝났습니다.\n", runner)
		wg.Done()
		return
	}

	fmt.Printf("%d 번째 주자가 %d 번째 주자에게 바통을 넘겼습니다.\n",
		runner, newRunner)
	baton <- newRunner
}
