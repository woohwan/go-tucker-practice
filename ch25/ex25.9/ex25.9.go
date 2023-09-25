// context with value
package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) // goroutine 하나 사용

	ctx := context.WithValue(context.Background(), "number", 9) // key-value 형식으로 전달
	go square(ctx)
	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square: %d", n*n)
	}
	wg.Done()
}
