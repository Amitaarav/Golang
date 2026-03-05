package main

import "fmt"

func main() {
    fmt.Println("Arrays in GoLang")

    var fruitList [4]string

    fruitList[0] = "Apple"
    fruitList[1] = "Tomato"
    fruitList[2] = "Banana"
    fruitList[3] = "Peach"

    fmt.Println("Fruit List is", fruitList)
    fmt.Println("Fruit List is", len(fruitList))
	fmt.Println("Fruit List is", fruitList[2])
}