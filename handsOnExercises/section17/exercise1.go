/*
create a value and assign it to a variable
print the address of that value
*/

package main

import "fmt"

func main() {
	value := "value"
	fmt.Println(value)
	fmt.Println(&value)
}
