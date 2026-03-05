package main

import (
	"fmt"
	"math"
)


const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	// var z uint = f
	fmt.Println(x, y, z)

	// Type inference
	v := "string"
	// v = 7
	fmt.Printf("Type %T\n", v)

	i := 42           // int
	f1 := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Println(i, f1, g)

	// constant
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	// Constants cannot be declared using the := syntax.
	// can not change the value
	// evaluated during compile time, not run tine

	// Predefined
	// number boolean string char

	// Numeric constant
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

}
