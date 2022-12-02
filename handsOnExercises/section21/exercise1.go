/*
In addition to the main goroutine, launch two additional goroutines
- each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exists
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go sayHi5()
	go sayBye10()
	wg.Wait()
	fmt.Println("greetings complete.")
}

func sayHi5() {
	for i := 0; i < 5; i++ {
		fmt.Println("hi!")
	}
	fmt.Println("I said hi five times.")
	wg.Done()
}

func sayBye10() {
	for i := 0; i < 10; i++ {
		fmt.Println("bye!")
	}
	fmt.Println("I said bye ten times.")
	wg.Done()
}
