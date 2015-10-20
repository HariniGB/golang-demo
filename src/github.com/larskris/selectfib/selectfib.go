package main

import "fmt"

func fibonacci(c, quit chan int) {
	fmt.Println("in fibornacci func")
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println("in go func()")
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fmt.Println("before fibonacci()")
	fibonacci(c, quit)
}