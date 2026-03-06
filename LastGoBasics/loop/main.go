package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string{
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64{
	if v:= math.Pow(x, n); v < lim {
		return v
	}else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// value of v cant be access outside the if block
	return lim
}
func main(){
	sum := 1
	for {
		if(sum > 100){
			break
		}
		sum += sum
	}
	fmt.Println(sum)
	fmt.Println(sqrt(2), sqrt(-4), sqrt(-3))
	fmt.Println(pow(3, 2, 10))

	// switch
	// for GOOS: run 'go tool dist list'
	switch os := runtime.GOOS; os{
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday{
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far days")
	}
}