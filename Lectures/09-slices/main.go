package main

import ("fmt"
"sort")

func main() {
    fmt.Println("Welcome to slices")
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)
	fmt.Println(fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:4]) //range are not inclusive of last index
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	highScores = append(highScores, 555, 666, 321) // memory allocation happens again
	fmt.Println("High Scores:", highScores)
	
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	//how to remove a value from slices based on index
	var course = []string{"reactjs","javascript","swift","python","ruby"}
	fmt.Println(course)
	var index int = 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)

}