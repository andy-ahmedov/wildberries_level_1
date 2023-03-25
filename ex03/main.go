package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	sum := 0
	var wg sync.WaitGroup

	for _, value := range numbers {
		wg.Add(1)
		go func(x int) {
			fmt.Println(x)
			sum += x * x
			defer wg.Done()
		}(value)
	}

	wg.Wait()
	fmt.Printf("Сумма квадратов чисел %v равна %d\n", numbers, sum)
}
