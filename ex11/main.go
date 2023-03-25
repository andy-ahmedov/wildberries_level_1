package main

import (
	"fmt"
)

func main() {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)
	interset := []int{}

	arr1 := []int{1, 5, 7, 4, 7, 6, 8, 2, 3, 9}
	arr2 := []int{3, 7, 9, 5, 7, 2, 9, 5, 6, 8}

	for i := 0; i < len(arr1); i++ {
		set1[arr1[i]] = true
		set2[arr2[i]] = true
	}
	fmt.Println()

	fmt.Println(set1)
	fmt.Println(set2)

	for i := 0; i < 10; i++ {
		if set1[i] && set2[i] {
			interset = append(interset, i)
		}
	}
	fmt.Println(interset)
}
