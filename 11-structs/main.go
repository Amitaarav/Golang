package main
import "fmt"
func main(){
	fmt.Println("Struct in golang")
	amit := User{"Amit", "amit@gmail.com", true, 22}
	fmt.Println(amit)
	fmt.Printf("Amit details are: %v\n", amit) //in case of struct %+v

}
type User struct {
	Name string
	Email string
	Status bool
	Age int
}