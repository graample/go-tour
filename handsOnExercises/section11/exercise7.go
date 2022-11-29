/*
create a slice of a slice of string([][]string)
store the following data in the slice:
- "James", "Bond", "Shaken, not stirred"
- "Miss", "Moneypenny", "Helloooooo, James."
range over the records, then range over the data in each record
*/

package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	z := [][]string{x, y}

	for i := range z {
		for j, w := range z[i] {
			fmt.Println(j, w)
		}
	}
}
