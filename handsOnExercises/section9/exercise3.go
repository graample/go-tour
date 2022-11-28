/*
Create a for loop using this syntax:

- for condition {}

Have it print out the years you have been alive.
*/

package main

import "fmt"

func main() {
	year := 1991
	for year <= 2022 {
		fmt.Println(year)
		year++
	}
}
