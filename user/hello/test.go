package main

import (
	"time"
	"fmt"
)

func main() {
	ch := make(chan int)
	go (func(ch chan int) {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(1 * 1e9)
		}
	})(ch)
	go (func(ch chan int) {
		for {
			fmt.Println(<-ch)
		}
	})(ch)
	time.Sleep(1e10)
}