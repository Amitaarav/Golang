package main

import (
	"fmt"
	"net/http"
	"io"
	"strings"

)

const url = "http://localhost:3000/"

// Function to start the server
func startServer() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, you've reached the local server!")
} 
func main(){
	fmt.Println("Hello World")
	go startServer()
	PerformGetRequest()
}

func PerformGetRequest(){
    const myUrl = "http://localhost:3000/get"

    response, err := http.Get(myUrl)

	if err != nil{
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	var responseString strings.Builder
	
	content,_ := io.ReadAll(response.Body) //content is in byte format
	byteCount,_:=responseString.Write(content)
	fmt.Println("Byte count is: ", byteCount)
	fmt.Println(string(content))
	
}