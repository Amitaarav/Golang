package main
import ("fmt"
	"bufio"
	"os"
)
func main(){
	welcome := "Welcome to my app"
	fmt.Println(welcome)

	// comma syntex
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")
	// comma ok syntex or error ok syntex
	input, err := reader.ReadString('\n') // read till '\n' is encountered
	fmt.Println("Thanks for rating, ", input)
	fmt.Println("Error is ", err);

	fmt.Printf("Type of this rating is %T", input)
}