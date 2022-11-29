/*
start with this slice:
- x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
append to that slice 52 and print
in ONE STATEMENT append to that slice these values and print:
53, 54, 55
append to the slice this slice and print:
y := []int{56, 57, 58, 59, 60}
*/

package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}
