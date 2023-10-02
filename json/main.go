package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string //Upper case bacause the json don't work
	Last  string
	Age   int
}

func main() {

	//Marshal
	p1 := person{
		First: "Luca",
		Last:  "Far",
		Age:   19,
	}

	p2 := person{
		First: "Riv",
		Last:  "DvD",
		Age:   78,
	}

	people := []person{p1, p2}
	fmt.Println(people)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error:",err)
	}
	fmt.Println(string(bs))

	//Unmarshal
	var people1 []person//create a slice for the output of the json
	json.Unmarshal(bs, &people1)//bs=json in byte, &people1 poit for the output of methods
	fmt.Println(people1)
}
