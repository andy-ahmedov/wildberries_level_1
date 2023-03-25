package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "aabcd"
	set := make(map[string]bool)

	for _, value := range str {
		lowercase := strings.ToLower(string(value))
		if set[lowercase] != true {
			set[lowercase] = true
		} else {
			fmt.Println("false")
			return
		}
	}
	fmt.Println("true")
}
