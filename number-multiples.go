package main

import (
	"fmt"
	"time"
)

func multiplesOfTwo(ch chan int) {
	for i := 2; ; i += 2 {
		ch <- i
		time.Sleep(time.Millisecond * 100)
	}
}

func multiplesOfThree(ch chan int) {
	for i := 3; ; i += 3 {
		ch <- i
		time.Sleep(time.Millisecond * 150)
	}
}

func main() {
	ch2 := make(chan int)
	ch3 := make(chan int)

	go multiplesOfTwo(ch2)
	go multiplesOfThree(ch3)

	var num2, num3 int
	for {
		select {
		case num2 = <-ch2:
		case num3 = <-ch3:
		}
		if num2 == num3 {
			fmt.Println("Matching number:", num2)
		}
	}
}
