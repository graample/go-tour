package main

import (
	"fmt"

	"github.com/graample/go-tour/tree/main/handsOnExercises/section27/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
