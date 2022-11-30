/*
build and use an anonymous func
*/

package main

import "fmt"

func main() {
	func() {
		fmt.Println("anonymous func used")
	}()
}
