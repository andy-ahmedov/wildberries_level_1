package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)

	for _, value := range sequence {
		set[value] = true
	}

	fmt.Println(set)
}
