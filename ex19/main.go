package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Println("Enter the string")
	fmt.Scan(&str)

	tmp := []rune(str)
	var tmp2 []rune

	for i := len(tmp) - 1; i >= 0; i-- {
		tmp2 = append(tmp2, tmp[i])
	}

	for _, value := range tmp2 {
		fmt.Printf("%c", value)
	}
	fmt.Println()
}
