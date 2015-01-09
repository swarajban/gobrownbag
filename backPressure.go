// Simple backpressure example by creating a pipeline with a
// buffered channel
package main

import (
	"fmt"
	"time"
)

func main() {

	numElements := 10

	// Buffered channel for backpressure!
	supplyChannel := make(chan int, 5)
	getSupply(supplyChannel, numElements)

	addedOne := make(chan int)
	addOne(supplyChannel, addedOne)

	// print output
	for output := range addedOne {
		fmt.Printf("output: %v\n", output)
	}
}

// Supplies the supplyChannel with a stream
// of ascending integers up to n
func getSupply(supplyChannel chan int, n int) {
	go func() {
		for i := 1; i <= n; i++ {
			// does not block until channel is full
			supplyChannel <- i
			fmt.Printf("adding %v to supply\n", i)
		}
		close(supplyChannel)
	}()
}

// Adds on to each int in the input stream and
// places it on the output stream
func addOne(inputStream chan int, outputStream chan int) {
	go func() {
		for input := range inputStream {
			outputStream <- input + 1
			time.Sleep(time.Second)
		}
		close(outputStream)
	}()
}
