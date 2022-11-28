/*
create your own type "person" which will have an underlying type of "struct" so that it can store the following data:
- first name
- last name
- favorite ice cream flavors
create two VALUES of TYPE person. print out the values, ranging over the elements in the slice which stores the favorite flavors
*/

package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first:      "Bob",
		last:       "Bobberson",
		favFlavors: []string{"Peach", "Rocky Road"},
	}
	p2 := person{
		first:      "Doug",
		last:       "Douggerson",
		favFlavors: []string{"Cherry", "Vanilla"},
	}
	fmt.Println(p1, p2)
}
