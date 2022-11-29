/*
using a composite literal:
- create a slice of type int
- assign 10 values
use SLICING to create the following new slices,
which are then printed:
- [5, 10, 15, 20, 25]
- [30, 35, 40, 45, 50]
- [15, 20, 25, 30, 35]
- [10, 15, 20, 25, 30]
*/

package main

import "fmt"

func main() {
	x := []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}

	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}
