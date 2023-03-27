package main

import (
	"fmt"
)

func main() {
	p2 := main.NewPoint(2, 3)
	p1 := main.NewPoint(1, 3)
	dist := Distance(p1, p2)
	fmt.Println(dist)

}
