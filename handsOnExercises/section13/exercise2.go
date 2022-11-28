/*
take the code from the previous exercise, then store the values of type person in a map with the key of last name
access each value in the map
print out the values, ranging over the slice
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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	// fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}
}
