package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	// launch function to fill c1 with values
	go populate(c1)

	// launch function to dump values from c1 into c2
	go fanOutIn(c1, c2)

	// print every value in c2
	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// take every number from 1-100, place them into c
func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup

	// for every value in c1
	for v := range c1 {
		// create a waitgroup
		wg.Add(1)

		// launch a new goroutine to randomize value, then place into c2
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

// takes an integer and randomizes it
func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
