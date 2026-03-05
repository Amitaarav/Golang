package main

import "fmt"

func main(){
	fmt.Println("Structs in Golang")
	amit := User{"Amit", "amit@gmail.com", false, 22}
	fmt.Println(amit)
	fmt.Printf("Amit details are: %+v\n", amit)
	fmt.Printf("Amit details are: %+v\n", amit.Name)
	amit.GetStatus()
	amit.GetAge()
	amit.NewMail()
}

type User struct {
	Name  string
	Email string
	Status bool
	Age int
}

func (u User) GetStatus(){
	fmt.Println("Is user active: ", u.Status)
}

func (u User) GetAge(){
	fmt.Println("User is: ", u.Age, "years old")
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}