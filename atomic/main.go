package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)

	fmt.Println("프로그램 종료!")
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
}

func doWork(name string) {
	defer wg.Done()
	for {
		fmt.Printf("작업 진행 중: %s\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("작업을 종료합니다: %s\n", name)
			break
		}
	}
}
