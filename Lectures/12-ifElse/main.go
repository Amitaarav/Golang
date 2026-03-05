package main

import "fmt"

func main() {
	// if else
	if 10 > 5 {
		fmt.Println("10 is greater than 5")
	} else {
		fmt.Println("10 is not greater than 5")
	}

	// if else if else
	if 10 > 5 {
		fmt.Println("10 is greater than 5")
	} else if 10 < 5 {
		fmt.Println("10 is less than 5")
	}

	if 9%2 == 0 {
		fmt.Println("9 is even")
	}else{
		fmt.Println("9 is odd")
	}

	if num := 10; num > 5 {
		fmt.Println("10 is greater than 5")
	}else {
		fmt.Println("10 is not greater than 5")
	}
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
}