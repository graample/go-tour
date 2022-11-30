/*
assign a func to a variable, then call that func
*/

package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("assigned func used")
	}

	a()
}
