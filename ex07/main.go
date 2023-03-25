package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	var n int
	m := make(map[int]int)
	fmt.Println("In how many seconds should the program terminate?")
	fmt.Scan(&n)
	sigChannel := make(chan os.Signal, 1)
	keyChannel := make(chan int)
	var mutex sync.Mutex

	for i := 0; i < n; i++ {
		go func(i int, mute *sync.Mutex) {
			for {
				mute.Lock()
				key := <-keyChannel
				m[key] = key * 2
				fmt.Printf("\033[92mm\033[31m[%2d]\033[0m%5d | ", key, m[key])
				mute.Unlock()
			}
		}(i, &mutex)
	}

	for i := 0; ; i++ {
		select {
		case sig := <-sigChannel:
			close(keyChannel)
			fmt.Printf("\nProgram with %d workers terminated because a '%v' signal was detected\n", n, sig)
			os.Exit(0)
		default:
			if i%6 == 0 {
				time.Sleep(time.Second)
				fmt.Println()
			}
			keyChannel <- i
		}
	}
}
