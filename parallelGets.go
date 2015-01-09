package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// sleep a random # of milliseconds and respond on channel
func query(q string, respChan chan string) {
	time.Sleep(time.Duration(randInt(100, 500)) * time.Millisecond)
	respChan <- fmt.Sprintf("response for %v", q)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// fire off 5 async to query("x") and concatenate all responses
	numQueries := 5
	finalResponse := make([]string, numQueries)

	responseChannel := make(chan string)

	// make N queries that will respond on responseChannel
	for i := 0; i < numQueries; i++ {
		go func(index int) {
			query(strconv.Itoa(index), responseChannel)
		}(i)
	}

	// concatenate responses
	for i := 0; i < numQueries; i++ {
		resp := <-responseChannel
		fmt.Printf("Received response: %v\n", resp)
		finalResponse = append(finalResponse, resp)
	}

	// response is an array of strings
	fmt.Println(finalResponse)
	// fmt.Println("done")
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}
