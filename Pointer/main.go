package main

import "fmt"


//func pass by value and reference |f3| |f4|

//|f3|
func ByValue(x int){//assume the value
	x=x+5
	fmt.Println("Inside ByValue function x:",x)
}
//|f4|
func ByReference(x *int){ //Point to value
	*x=*x+5
	fmt.Println("Inside ByReference function x:",*x)}


func main() {
	//Pointer
	
	//*il valore dell'indirizzo, & l'indirizzo del valore

	i := 10
	p := &i
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(i)
	i = 20
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(i)

	//Pass by value and by reference with |f3| and |f4| function

	x:=2
	fmt.Println("Init x:",x)
	ByValue(x)
	fmt.Println("After ByValue function x:",x)
	ByReference(&x)
	fmt.Println("After ByReference function x:",x)

	/*
	Reference type:
		Pointers
		Slices
		Maps
		Channels
		Functions
		Interfaces

	Mutability type:
		slices
		maps
		pointers
	*/

	//metho

	d1:= dog{first: "Doggy",}
	d2:= &dog{first: "Henry",}
	d1.walk()
	d2.run()
	d1.walk()
	d2.run()

	//youngRun(d1)   //don't work because T -> methods t T
	//youngRun(d2)   //work because *T -> methods t T and t *T

}
type dog struct{
	first string
}

func (d dog) walk(){
	fmt.Println("I'm",d.first,"and I'm walking")
}

func (d *dog) run(){
	fmt.Println("I'm",d.first,"and I'm walking")
}
type yungin interface{
	walk()
	run()
}
func youngRun(y yungin){
	y.run()
}