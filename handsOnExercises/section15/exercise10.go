/*
create a func which encloses the scope of a variable
*/

package main

import "fmt"

func main() {
	a := foo(2)
	fmt.Println(a)

	b := foo(3)
	fmt.Println(b)
}

func foo(num int) int {
	fmt.Println("a new function has been ran!")
	x := 1
	fmt.Println("the value of x in this function is still", x)
	a := x + num
	fmt.Println("x and num have been added together:", a)
	return a
}
