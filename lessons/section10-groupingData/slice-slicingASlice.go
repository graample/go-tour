package main

import "fmt"

func main() {
	// declare an empty array of five integers
	x := []int{4, 5, 7, 8, 42}

	// print item at index 1
	fmt.Println(x[1])

	// print x
	fmt.Println(x)

	// print starting from index 1
	fmt.Println(x[1:])

	// print starting from index 1 and up to 3
	fmt.Println(x[1:3])

	// print index and value of all items in x
	for i, v := range x {
		fmt.Println(i, v)
	}

	// print index and value of all items in x
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
