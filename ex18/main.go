package main

import (
	"fmt"
	"sync"
)

type IncrementStruct struct {
	i int
	sync.Mutex
}

func (i *IncrementStruct) Increment() {
	i.Lock()
	i.i++
	i.Unlock()
}

func main() {
	inc := IncrementStruct{
		i: 0,
	}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			for j := 0; j < 20; j++ {
				inc.Increment()
			}
			defer wg.Done()
		}(&wg)
	}
	wg.Wait()
	fmt.Println(inc.i)
}
