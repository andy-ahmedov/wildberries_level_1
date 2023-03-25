package main

import "fmt"

func removeElement(array []int, index int) []int {
	slice := append(array[:index], array[index+1:]...)
	return slice
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice)
	slice = removeElement(slice, 4)
	fmt.Println(slice)
}
