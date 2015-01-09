package main

import (
	"fmt"
)

func seriesGenerator(n int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func factorial(x int) int {
	fact := 1
	series := seriesGenerator(x)
	for n := range series {
		fact = fact * n
	}
	return fact
}

func main() {
	f5 := factorial(5)
	fmt.Println(f5)
}
