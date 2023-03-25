package main

import (
	"fmt"
	"math/big"
)

func main() {
	str := "сегодня мне придется это оставить"

	size := 0
	var tmp2 []rune
	tmp := []rune(str)
	words := make(map[int][]rune)

	for i := 0; i < len(tmp); i++ {
		if i == len(tmp)-1 {
			tmp2 = append(tmp2, tmp[i])
		}
		if tmp[i] == ' ' || i == len(tmp)-1 {
			words[size] = tmp2
			tmp2 = []rune("")
			size++
		} else {
			tmp2 = append(tmp2, tmp[i])
		}
	}

	for i := size; i >= 0; i-- {
		tmp2 = append(tmp2, words[i]...)
		if i < size && i > 0 {
			tmp2 = append(tmp2, ' ')
		}
	}
	str2 := string(tmp2)
	fmt.Println(str2)
	a := big.NewInt(1<<55 + 1<<32 + 1<<27)
	fmt.Println(a)
}