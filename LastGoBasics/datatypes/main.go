package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	str    string     = "Hello"
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %q\n",str, str) // %s -> string , %q --> with double quote
}

// 1. basic
//  number, string, bool

// 2. Aggregate 
//  array(collection of similar type fo data types variable), struct(collection of hetrogenous data types)

// 3. Reference 
//  Pointers, slices, function, chanel, maps

// 4. Interface types

// int8 -127 to 128 (signed integer)
// uint8 0 to 255 (unsigned integer)
// always go with system defined architecture

// float32 float64 does not have system defined architecture

