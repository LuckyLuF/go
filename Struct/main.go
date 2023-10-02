package main

import "fmt"

type person struct{
	first string 
	last string
	age int
}
type agent struct{
	person person
	working bool
}

func main(){
	//Struct

	
	e:= person{
		first:"James",
		last:"Bond",
		age:50,
	}
	p :=new(person)
	p.age=19
	p.last="Farinella"
	p.first="Luca"
	fmt.Printf("p: %v\n", p)
	fmt.Printf("e: %v\n", e)
	r:= agent{
		person: person{
			first:"James",
			last:"Bond",
			age:50,
	},
	working: true,
	}
	a :=new(agent)
	a.person.age=23
	a.person.first="John"
	a.person.last="Wick"
	a.working=false
	fmt.Printf("p: %v\n", p)
	fmt.Printf("e: %v\n", e)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("r: %v\n", r)

	q:= struct{ // anonimous struct
		first string 
		last string
		age int
	}{
		first: "Cane",
		last: "Dog",
		age: 89,
	}
	fmt.Printf("%T struct anionima tipo \n", q)
	fmt.Printf("%v struct anionima valori\n", q)
	q = e  //if they have the same field, q copy the e values
	fmt.Printf("%T struct anionima tipo\n", q)
	fmt.Printf("%v struct anionima valori\n", q)
	//q = a don't work because they have different field
	
}