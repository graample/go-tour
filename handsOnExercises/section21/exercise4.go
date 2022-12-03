/*
fix the race condition you created in the previous exercise by using a mutex
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(100)
	for x := 0; x < 100; x++ {
		go func() {
			mu.Lock()
			fmt.Println(counter)
			i := counter
			i++
			counter = i
			fmt.Println(counter)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end:", counter)
}
