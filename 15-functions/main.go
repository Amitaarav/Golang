package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeting := greet("Shubham")
	fmt.Println(greeting)
	greaterTwo()

	result := add(3, 5)
	fmt.Println("Result is :",result)

	proResult := proAdder(1,2,3,4,5,6,7,8,9,10)
	fmt.Println("Pro adder result is :",proResult)

}
func greet(name string) string {
	return "Hello " + name
}

func greaterTwo(){
	fmt.Println("Another method")
}
func add(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}