package main

import "fmt"

// jwtToken := 30000 this is allowed only in method scope // use it when declared otherwise it will give error

// capital letter is used for public variable

const LoginToken string = "token"
func main() {
    var username string  = "Amit" // use it when declared otherwise it will give error
    fmt.Println(username) //Println is used to print the output
	fmt.Printf("variable is of type : %T \n", username)  //Printf is used to print the output with the type of variable

    // implicit type
    var website = "amit.com"
    fmt.Println(website)

    //no var style
    numberOfUser := 300000
    fmt.Println(numberOfUser)

    fmt.Println(LoginToken)
    fmt.Printf("variable is of type : %T \n", LoginToken)
}