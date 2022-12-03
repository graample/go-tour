/*
fix the race condition you created in exercise #4 by using package atomic
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup
	wg.Add(100)
	for x := 0; x < 100; x++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end:", counter)
}
