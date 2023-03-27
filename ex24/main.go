package main

import (
	"fmt"
	"github.com/andy-ahmedov/wildberries_level_1/ex24/structure"
)

func main() {
	par := structure.NewPoint(2, 3)
	par2 := structure.NewPoint(5, 7)

	result := structure.Distance(par, par2)
	fmt.Println(result)
}
