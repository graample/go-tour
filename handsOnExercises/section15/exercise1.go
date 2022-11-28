/*
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results
*/

package main

import "fmt"

func main() {
	x := foo(6)
	y, z := bar(7, "eight")
	fmt.Println(x, y, z)
}

func foo(num int) int {
	return num
}

func bar(num int, str string) (int, string) {
	return num, str
}
