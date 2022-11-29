/*
using a composite literal:
- create an ARRAY which holds 5 values of type int
- assign values to each index position
range over the array and print the values out
using format printing:
- print out the TYPE of the array
*/

package main

import "fmt"

func main() {
	x := [5]int{5, 10, 15, 20, 25}

	for k, v := range x {
		fmt.Println(k, v)
	}

	fmt.Printf("%T\n", x)
}
