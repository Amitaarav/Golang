package main

import "fmt"

func add(x int, y int) int{
	return x + y
}

// a function can return multiple values 

func swap(x, y string)(string, string){
	return y, x
}

// func split(sum int)(x, y int){
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return  // naked return
// }

func split(sum int)(int, int){
	x := sum * 4 / 9
	y := sum - x
	return x, y
}

// var c, python, java bool
// A var statement can be at package or function level.

// var name type

var i, j int = 1, 2
func main(){
	fmt.Println(add(22,34))
	// fmt.Println(swap("hello", "world"))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))

	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}