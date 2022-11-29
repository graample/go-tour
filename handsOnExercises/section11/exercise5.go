/*
start with this slice:
- x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
use APPEND and SLICING to get the values here which you should then print:
- [42, 43, 44, 48, 49, 50, 51]
*/

package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(append(x[:3], x[6:]...))
}
