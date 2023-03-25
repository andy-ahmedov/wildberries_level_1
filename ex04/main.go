package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var n int
	color := 3
	fmt.Println("Enter the number of workers: ")
	fmt.Scan(&n)
	outChannel := make(chan int)
	sigChannel := make(chan os.Signal, 1)
	var outMutex sync.Mutex
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)

	for i := 0; i < n; i++ {
		go func(i int, outMute *sync.Mutex) {
			for {
				outMute.Lock()
				div := i / 7
				switch div % 2 {
				case 0:
					color = 3
				case 1:
					color = 9
				}
				fmt.Print("\033[", color*10+i%7+1, "m")
				switch n / 10 {
				case 0:
					fmt.Printf("(%d)\033[0m:%5d ", i+1, <-outChannel)
				case 1, 2, 3, 4, 5, 6, 7, 8, 9:
					fmt.Printf("(%2d)\033[0m:%5d ", i+1, <-outChannel)
				default:
					fmt.Printf("(%3d)\033[0m:%5d ", i+1, <-outChannel)
				}
				outMute.Unlock()
			}
		}(i, &outMutex)
	}

	for i := 0; ; i++ {
		select {
		case sig := <-sigChannel:
			close(outChannel)
			fmt.Printf("\nProgram with %d workers terminated because a '%v' signal was detected\n", n, sig)
			os.Exit(0)
		default:
			if i%6 == 0 {
				time.Sleep(time.Second)
				fmt.Println()
			}
			outChannel <- rand.Intn(1000 - 1 + 1)
		}
	}
}
