package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
    fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - Hellooo baki sab kaise hain"

	file,err := os.Create("./mylcogofile1.txt")
	if err != nil {
		panic(err) // shut down execution of the program
	}

	// for writting
	length,err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("length is: ", length)
	defer file.Close()

	readFile("./mylcogofile1.txt")


}

// read the file
func readFile(fileName string){
	databytes,err:= ioutil.ReadFile(fileName)
	checkNilError(err)
	fmt.Println("content of the file is: ", string(databytes))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}