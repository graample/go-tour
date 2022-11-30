/*
create a func with the identifier foo that:
- takes in a variadic parameter of type int
- pass in a value of type []int into your func (unfurl the []int)
- return the sum of all values of type int passed in

create a func with the identifier bar that:
- takes in a parameter of type []int
- return the sum of all values of type int passed in
*/

package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	fmt.Println(foo(x...))
	fmt.Println(bar(x))
}

func foo(arr ...int) int {
	// sum the array
	var total int
	for _, v := range arr {
		total += v
	}
	return total
}

func bar(arr []int) int {
	var total int
	for _, v := range arr {
		total += v
	}
	return total
}
