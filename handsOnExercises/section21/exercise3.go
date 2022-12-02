/*
using goroutines, create an incrementer program
	have a variable to hold the incrementer value
	launch a bunch of goroutines that should:
		read the incrementer value
			store it in a new variable
		yield the processor with Gosched()
		increment the new variable
		write the value in the new variable back to the incrementer variable
use waitgroups to wait for all of your goroutines to finish
the above will create a race condition
prove that it is a race condition by using the -race flag
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup
	wg.Add(100)
	for x := 0; x < 100; x++ {
		go func() {
			fmt.Println(counter)
			i := counter
			runtime.Gosched()
			i++
			counter = i
			fmt.Println(counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end:", counter)
}
