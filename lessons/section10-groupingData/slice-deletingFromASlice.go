package main

import "fmt"

func main() {
	// declare an empty array of five integers and print
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	// append values into x and print
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	// append y values into x and print
	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

	// delete items at index 2 and 3
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
