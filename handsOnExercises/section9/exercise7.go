/*
Create a program that uses "else if" and "else"
*/

package main

import "fmt"

func main() {
	x := ""
	if x == "hi" {
		fmt.Println("hello to you too!")
	} else if x == "bye" {
		fmt.Println("see ya!")
	} else {
		fmt.Println("you are being ignored")
	}
}
