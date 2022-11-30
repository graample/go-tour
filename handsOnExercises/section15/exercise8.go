/*
create a func which returns a func
assign the returned func to a variable
call the returned func
*/

package main

import "fmt"

func main() {
	a := func() func() {
		return func() {
			fmt.Println("returned func used")
		}
	}
	b := a()
	b()
	c := outer()
	fmt.Println(c())
}

func outer() func() string {
	return func() string {
		return "inner text"
	}
}
