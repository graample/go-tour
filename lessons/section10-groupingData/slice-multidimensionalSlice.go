package main

import "fmt"

func main() {
	// declare two string arrays
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	// create a multidimensional slice from the two arrays
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
