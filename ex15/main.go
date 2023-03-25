package main

import "fmt"

var justString string

func createHugeString(n int) string {
	str := make([]rune, n)
	for i := range str {
		str[i] = 'ё'
	}
	return string(str)
}

func someFunc() string {
	v := createHugeString(1 << 10)
	if len(v) < 100 {
		justString = v[:100]
	} else {
		justString = v
	}
	return justString
}

func main() {
	str := someFunc()
	fmt.Println(str, len([]rune(str)))
}
