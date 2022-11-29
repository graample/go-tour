/*
Using the code from the previous example,
add a record to your map.
now print the map out using the "range" loop
*/

package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	x["yee_raymond"] = []string{"JavaScript", "React", "node"}
	// fmt.Println(x)

	for name, list := range x {
		fmt.Println(name)
		for i, thing := range list {
			fmt.Println(i, thing)
		}
	}
}
