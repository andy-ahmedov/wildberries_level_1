package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cansel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine termination signal received")
				return
			default:
				fmt.Println("Goroutine works...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second * 5)
	cansel()
	time.Sleep(time.Second)
	fmt.Println("Program completed")
}
