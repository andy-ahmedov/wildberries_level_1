package main

import (
	"fmt"
	"time"
)

func ft_sleep(d time.Duration) {
	timer := time.NewTimer(d)
	<-timer.C
}

func main() {
	ft_sleep(time.Second * 3)
	fmt.Println("Times up!")
}
