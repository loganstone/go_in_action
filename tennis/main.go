package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)
	wg.Add(2)

	go player("Logan", court)
	go player("Hyunsuk", court)

	court <- 0

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("%s 선수가 승리 했습니다.\n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("%s 선수가 공을 받아치지 못했습니다.\n", name)
			close(court)
			return
		}

		if ball == 0 {
			fmt.Printf("%s 선수가 서브를 합니다.\n", name)
		} else {
			fmt.Printf("%s 선수가 %d 번째 공을 받아쳤습니다.\n", name, ball)
		}

		ball++
		court <- ball
	}
}
