package main

import "fmt"

func main() {
	// declare an empty array of five integers
	var x [5]int
	fmt.Println(x)

	// insert 42 at the fourth position of the x array
	x[3] = 42
	fmt.Println(x)

	// print the length
	fmt.Println(len(x))
}
