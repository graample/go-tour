/*
Using the code from the previous example,
delete a record from your map.
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

	delete(x, `bond_james`)

	for name, list := range x {
		fmt.Println(name)
		for i, thing := range list {
			fmt.Println(i, thing)
		}
	}
}
