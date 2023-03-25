package main

import "fmt"

func main() {
	mas := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	downstream := make(chan int)
	upstream := make(chan int)
	go func() {
		for _, value := range mas {
			downstream <- value
		}
		close(downstream)
	}()

	go func(upstream chan int, downstream chan int) {
		for value := range upstream {
			downstream <- value * 2
		}
		close(downstream)
	}(downstream, upstream)

	for value := range upstream {
		fmt.Println(value)
	}
}
