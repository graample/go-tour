/*
Print out the remainder(modulus) which is found for each number between 10 and 100 when it is divided by 4.
*/
package main

import "fmt"

func main() {
	for i := 11; i < 100; i++ {
		fmt.Println(i, i%4)
		// fmt.Printf("When %v is divided by 4, the remainder (aka modulus) is %v\n", i, i%4)
	}
}
