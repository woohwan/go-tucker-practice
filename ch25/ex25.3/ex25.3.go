/*
채널에서 대기: range chan
main은 wg.Wait(), goroutine은 range chan에서 무한 대기 --> goroutine leak 발생
해결 방법: 데이터를 다 송신 후 channel close()
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	// go routine과 통신할 채널 생성, 크기 2
	ch := make(chan int, 2)
	// 통신할 go routine 생성
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2 // 데이터를 송신
	}
	close(ch) // 채널 차단 -> 정상 종료

	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	// 무한 대기 발생
	for n := range ch { // channel이 닫히면 loop exit
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}
