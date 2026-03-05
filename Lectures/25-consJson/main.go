package main

import (
	"fmt"
	"encoding/json"
)

type course struct {
	Name   		string `json:"coursename"`
	Price 		int 
	Platform 	string `json:"website"`
	Password 	string `json:"-"`
	Tags 		[]string `json:"tags,omitempty"` // omitempty will not include empty tags in json
}

func main(){
	fmt.Println("Welcome to the JSON tutorial")
	DecodingJSON()
}

func DecodingJSON(){

	jsonDataFromWeb := []byte(`
	{
                "Name": "Go Bootcamp",
                "Price": 399,
                "Platform": "LearnCodingOnline.in",   
                "Password": "pqr123",
                "Tags": [
                        "back-end",
                        "go"
                ]
        }
	`)

	var lcoCourses course

	checkValid := json.Valid(jsonDataFromWeb)
	if(checkValid){
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
	}else{
		fmt.Println("Invalid JSON")
	}
}