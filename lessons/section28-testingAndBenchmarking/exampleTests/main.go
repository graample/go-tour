package main

import (
	"fmt"

	"github.com/graample/go-tour/tree/main/lessons/section28-testingAndBenchmarking/exampleTests/acdc"
)

func main() {
	fmt.Println(acdc.Sum(2, 3))
	fmt.Println(acdc.Sum(2, 3, 4, 5, 6, 7, 8, 9))
}
