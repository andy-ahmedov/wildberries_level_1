package main

import "fmt"

func main() {
	num1 := 12
	num2 := 4

	fmt.Println(num1, num2)

	num1, num2 = num2, num1

	fmt.Println(num1, num2)
}
