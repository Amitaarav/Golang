package main

import "fmt"

// jwtToken := 30000 this is allowed only in method scope 
// use it when declared otherwise it will give error

// capital letter is used for public variable

// var name type
const LoginToken string = "token"

// k := 9 // use var before name since it is out side the function
func main() {
    var username string  = "Amit" // use it when declared otherwise it will give error
    fmt.Println(username) //Println is used to print the output
	fmt.Printf("variable is of type : %T \n", username)  //Printf is used to print the output with the type of variable

    // implicit type
    var website = "amit.com"
    fmt.Println(website)

    //no var style
    // we can do inside function not outside the function
    // only accessible inside the function
    numberOfUser := 300000
    fmt.Println(numberOfUser)

    fmt.Println(LoginToken)
    fmt.Printf("variable is of type : %T \n", LoginToken)
}

// If initial value of the variable does not matter --> var i int
// If we know initial value of the variable then we can go with i := 1 pattern
