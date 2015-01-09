// Merge 3 'feeds' that produce values randomly every 500ms -> 2s
// into a single response
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Merge three feeds with a single select statement and no lockcs
func mergeFeeds(feedA, feedB, feedC <-chan string) <-chan string {
	merged := make(chan string)
	go func() {
		for {
			// select blocks until one feed has a response,
			select {
			case resp := <-feedA:
				merged <- resp

			case resp := <-feedB:
				merged <- resp

			case resp := <-feedC:
				merged <- resp
			}
		}
	}()
	return merged
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// create 3 new 'feeds'
	feedA := getFeed("A")
	feedB := getFeed("B")
	feedC := getFeed("C")

	// merged channel fans-in all feeds
	merged := mergeFeeds(feedA, feedB, feedC)

	// read merged responses forever
	for {
		resp := <-merged
		fmt.Printf("Received response: %v", resp)
	}

	// never returns
}

func getFeed(name string) <-chan string {
	c := make(chan string)
	entryCount := 1
	go func() {
		for {
			time.Sleep(time.Duration(randInt(500, 2000)) * time.Millisecond)
			c <- fmt.Sprintf("Entry %v from feed %v\n", entryCount, name)
			entryCount += 1
		}

	}()
	return c
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}
