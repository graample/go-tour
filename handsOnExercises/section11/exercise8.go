/*
create a map of TYPE string
which is a person's "last_first" name,
and a value of type []string
which stores their favorite things.
store three records in your map.
print out all of the values,
along with their index position in the slice.

`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
`no_dr`, `Being evil`, `Ice cream`, `Sunsets`
*/

package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	// fmt.Println(x)

	for name, list := range x {
		fmt.Println(name)
		for i, thing := range list {
			fmt.Println(i, thing)
		}
	}
}
