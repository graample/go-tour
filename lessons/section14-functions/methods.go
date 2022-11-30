package main

import "fmt"

// definition for type person
type person struct {
	first string
	last  string
}

// secretAgents are also persons, but with a license to kill
type secretAgent struct {
	person
	ltk bool
}

// defining method speak for type secretAgent
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

// using it in practice
func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak()
}
