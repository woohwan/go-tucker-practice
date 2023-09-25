// deadlock을 피하기 위한 작업영역 분리

package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Do()
}

// 작업영역 분리: Job 분리 -> 영역분리 -> 공유자원 X
type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index) // 각각의 작업
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index*j.index)
}

func main() {
	var JobList [10]Job

	for i := 0; i < 10; i++ { // 10가지 작업을 나누어 할당
		JobList[i] = &SquareJob{i}
	}

	wg := sync.WaitGroup{}
	wg.Add(10)

	// for i := 0; i < 10; i++ {
	// 	job := JobList[i] // 각 각업을 고루틴으로 실행
	// 	go func() {
	// 		job.Do()
	// 		wg.Done()
	// 	}()

	// }
	for _, job := range JobList {
		job := job
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}
