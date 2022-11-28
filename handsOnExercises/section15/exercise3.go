/*
use the "defer" keyword to show that a deferred func runs after the func containing it exits.
*/

package main

import "fmt"

func main() {
	defer inner()
	outer()
}

func inner() {
	fmt.Println("this should run last")
}

func outer() {
	fmt.Println("this should run first")
}
