package main

import "fmt"
func main() {
    fmt.Println("Loops in Go")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	for d:=0; d < len(days); d++ {
		fmt.Println(days[d]) // no ++d
	}
	for i,day:= range days {
		fmt.Printf("Day: %v and index: %v\n",day,i)
	}

	roughValue := 1
	for roughValue < 10 {
		if roughValue == 2 {
		    goto lco
		}
		if roughValue == 5 {
			roughValue++
		    continue
		}
		fmt.Println("Value: ", roughValue)
		roughValue++
	}
	lco:
	fmt.Println("Learning to code online")
}