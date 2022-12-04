/*
fix this code using:
- a func literal, aka, anonymous self-executing func
- a buffered channel
*/
package main

import (
	"fmt"
)

func main() {
	funcLiteral()
	buffered()
}

// func literal
func funcLiteral() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

// buffered channel
func buffered() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
