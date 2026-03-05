package main // package declairation of package main
// usecases
// 1. 

// package name have to be short,small letter

import (
	"fmt"
	"math"
	"math/rand"
	) // import from standard library/ external packages

func main() {
	fmt.Println("My favorite number is", rand.Intn(10)) // UTF-8 encoding supports - 120k characters
	fmt.Println("Value of Pi: ",math.Pi)
	// ascii supports 127 characters
}
// P is capital -- 
