package main

import (
	"fmt"
)

func main() {
	p1 := NewPoint(5, 7)
	p2 := NewPoint(5, 7)

	dist := Distance(p1, p2)
	fmt.Println(dist)

}
