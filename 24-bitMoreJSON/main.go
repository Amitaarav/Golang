package main

import (
	"fmt"
	"encoding/json"
	
)
// CONVERTING DATA IN VALID JSON FORMAT
type course struct {
	Name   		string `json:"coursename"`
	Price 		int 
	Platform 	string `json:"website"`
	Password 	string `json:"-"`
	Tags 		[]string `json:"tags,omitempty"` // omitempty will not include empty tags in json
}

func main(){
	fmt.Println("Welcome to json video")
	EncodeJson()
}

func EncodeJson(){
	lcoCourses := []course{
		{"ReactJs Bootcamp",299,"LearnCodingOnline.in","abc123",[]string{"web-dev","js"}},
		{"MERN Bootcamp",199,"LearnCodingOnline.in","xyz123",[]string{"full-stack","js"}},
		{"Go Bootcamp",399,"LearnCodingOnline.in","pqr123",[]string{"back-end","go"}},
	}
	// package this data as json data
	finalJson, err  := json.MarshalIndent(lcoCourses, "", "\t") // implementation of json.Marshal is done by the json package
	if(err != nil){	
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n",finalJson)
}
