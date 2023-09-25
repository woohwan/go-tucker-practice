package main

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan int) // 크기 0인 채널 생성
	ch := make(chan int, 2) // 보관함이 있으시, 곧바로 exit --> Never Print 출력

	go square()
	// 데이터를 넣었는데, 가져가는 루틴이 없어서 무한 대기. 원래는 main에서 wait가 없어서 곧바로 exit
	// 해서 sleep 이 출력되면 안되지만,  출력 됨.
	ch <- 9
	fmt.Println("Never print")
}

func square() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("sleep")
	}
}
