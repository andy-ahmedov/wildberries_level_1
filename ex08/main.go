package main

import "fmt"

func print_binary(num int64) {
	for i := 63; i >= 0; i-- {
		if num&(1<<i) != 0 {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
		if i%4 == 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func main() {
	var num int64
	var res int64
	var i_bit int
	var bit int
	fmt.Println("Choose which bit you want to change")
	fmt.Scan(&i_bit)
	fmt.Println("What do you want to record?")
	fmt.Scan(&bit)

	num = 9223372036854775807

	print_binary(num)

	mask := int64(1) << int64(i_bit-1)

	print_binary(mask)
	switch bit {
	case 1:
		res = num | mask
	case 0:
		res = num & ^mask
	}

	fmt.Println()
	num = res

	print_binary(num)

}
