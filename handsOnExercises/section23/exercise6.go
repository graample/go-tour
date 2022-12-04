/*
write a program that
- adds 100 numbers to a channel
- pull the numbers off and print them
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	send(c)
	receive(c)
	fmt.Println("done")
}

func send(c chan int) {
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
