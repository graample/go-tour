/*
using a composite literal:
- create a slice of type int
- assign 10 values
range over the slice and print the values out
using format printing:
- print out the TYPE of the slice
*/

package main

import "fmt"

func main() {
	x := []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}

	for k, v := range x {
		fmt.Println(k, v)
	}

	fmt.Printf("%T\n", x)
}
