package main

import (
	"fmt"
)

func main() {
	// a map is an object lol
	// create map of KEY TYPE string and VALUE TYPE int
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)
	fmt.Println("this is how old james is:", m["James"])

	// if the key isn't in the map, zero value is returned
	fmt.Println(m["Barnabas"])

	// run a check to see if it exists
	if v, ok := m["Barnabas"]; ok {
		fmt.Println("This key exists!", v)
	} else {
		fmt.Println("This key doesn't exist!")
	}
}
