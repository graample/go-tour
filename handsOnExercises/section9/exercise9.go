/*
Create a program that uses a switch statement
with the switch expression specified as a variable of TYPE string
with the IDENTIFIER "favSport"
*/

package main

import "fmt"

func main() {
	favSport := "sportsball"
	switch favSport {
	case "basketball", "skiing", "swimming", "surfing":
		fmt.Println("false")
	case "sportsball":
		fmt.Println("the best sport ever invented!")
	}
}
