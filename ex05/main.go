package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var n int
	fmt.Println("In how many seconds should the program terminate?")
	fmt.Scan(&n)
	ch := make(chan int)
	timer := time.NewTimer(time.Duration(n) * time.Second)

	go func() {
		for i := 0; ; i++ {
			fmt.Printf("%3d ", <-ch)
		}
	}()

	for i := 0; ; i++ {
		select {
		case <-timer.C:
			close(ch)
			fmt.Println("\n\t\t\033[91mTime is up!")
			os.Exit(0)
		default:
			if i%10 == 0 {
				fmt.Println()
			}
			ch <- rand.Intn(1000)
		}
	}
}
