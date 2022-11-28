/*
create and use an anonymous struct.
*/

package main

import "fmt"

func main() {
	bub := struct {
		name  string
		claws bool
	}{
		name:  "Wolverine",
		claws: true,
	}
	fmt.Println(bub.name)
	fmt.Println("Does he have claws?", bub.claws)
}
