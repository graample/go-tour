package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	// start the send function
	go send(even, odd)

	// start the receive function
	go receive(even, odd, fanin)

	// for all values in the fanin channel
	for v := range fanin {
		// print the values
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// send channel
func send(even, odd chan<- int) {
	// out of 10
	for i := 0; i < 10; i++ {
		// put even and odd numbers in their respective channels
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	// close channels
	close(even)
	close(odd)
}

// receive channel (takes values from odd/even and places them into fanin)
func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	// launch goroutine to receive values from even
	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	// launch goroutine to receive values from odd
	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
