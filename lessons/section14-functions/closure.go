package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()

	// x will be a local variable inside of a
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	// x will be a different variable inside of b
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
