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

	// delete an entry
	delete(m, "James")
	fmt.Println(m)

	// you won't get an error for keys that don't exist
	delete(m, "Ian Fleming")
	fmt.Println(m)

	// if item is found, print it then delete it
	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss Moneypenny")
	}

	fmt.Println(m)
}
