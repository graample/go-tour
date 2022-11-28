/*
create a person struct
create a func called "changeMe" with a *person as a parameter
	change the value stored at the *person address
create a value of type person
	print out the value
in func main
	call "changeMe"
in func main
	print out the value
*/

package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p1 := person{
		firstName: "Barry",
		lastName:  "Allen",
		age:       0,
	}
	fmt.Println("before", p1)
	changeMe(&p1)
	fmt.Println("after", p1)
}

func changeMe(p *person) {
	// p.firstName = "Big"
	(*p).firstName = "Little"
	p.lastName = "Chungus"
	p.age = 9999
}
