package main

import (
	"fmt"
	"math/rand"
)

func partition(first int, size int, arr []int) int {
	pivot := arr[size]
	tmp := first

	for i := first; i < size; i++ {
		if arr[i] < pivot {
			arr[i], arr[tmp] = arr[tmp], arr[i]
			tmp++
		}
	}
	arr[tmp], arr[size] = arr[size], arr[tmp]
	return tmp
}

func quickSort(first int, size int, arr []int) {
	if first < size {
		part := partition(first, size, arr)
		quickSort(first, part-1, arr)
		quickSort(part+1, size, arr)
	}
}

func main() {
	array := []int{}

	for i := 0; i < 25; i++ {
		array = append(array, rand.Intn(100))
	}
	fmt.Println(array)
	quickSort(0, len(array)-1, array)
	fmt.Println(array)
}
