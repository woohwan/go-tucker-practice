// Context 관련, context는 작업지시서 역할
// goroutine의 생성, 취소, 시간, 값 전달 등
// 5 초후에 goroutine 작업  cancel. 시그널은 conext.Done()을 통해서 이루어 짐.
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background()) // Backgroud context는 비어있는 default context
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel() //canncel goroutine
	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second) // 1초마 한 번씩 signal을 보내는 channel
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}
