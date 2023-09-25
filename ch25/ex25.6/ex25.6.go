/*
select: 여러 channel에서 동시에 데이터를 기다릴 때 사용. --> 한 번만 사용 가능
그래서, 무한루프와 (for {}) 병행해서 사용

time package의
- Tick()은 일정간격으로 신호를 주는 채널을 반환  --> 일정 간격으로 실행 시
- After()는 일정 시간 대기 후 한번 만 신호를 주는 채널 반환
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	// create channel
	ch := make(chan int)

	go square(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second) // 1초마다 signal(data)을 주는 channel 생성
	terminated := time.After(10 * time.Second)

	for {
		select { // 동시에 2가지 이상 조건을 만족할 시 select가 random하게 선택함.
		case <-tick:
			fmt.Println("Tick")
		case <-terminated:
			fmt.Println("Terminated")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}
