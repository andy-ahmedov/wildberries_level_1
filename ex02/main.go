package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}

	var w sync.WaitGroup
	for _, val := range numbers {
		w.Add(1)
		go func() {
			fmt.Println(val * val)
			defer w.Done()
		}()
		w.Wait()
	}
}
