/*
create a type person struct
attach a method speak to type person using a pointer receiver
- *person
create a type human interface
- to implicitly implement the interface, a human must have the speak method
create func "saySomething"
- have it take in a human as a parameter
- have it call the speak method
show the following in your code
- you CAN pass type *person into saySomething
- you CANNOT pass type person into saySomething
*/

package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func (p *person) speak() {
	fmt.Println(*&p.firstName, *&p.lastName, "says: bark!")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		firstName: "bob",
		lastName:  "bobberson",
	}
	saySomething(&p1)
}
