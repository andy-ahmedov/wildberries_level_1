package main

import "fmt"

func main() {
	set := make(map[int][]float32)
	arr := [10]float32{10.5, 13.2, 6.0, -24.4, 21.4, 34.5, 23.6, 6.7, 5.7, 17.8}

	for i := 0; i < 10; i++ {
		key := int(arr[i]) / 10 * 10
		set[key] = append(set[key], arr[i])
	}
	fmt.Println(set)
}
