package main

import (
	"fmt"
	"log"
	"strconv"
)

//|f5|   variadic paramaters 
func call(names ...string){ // ... -> the function have multiple parameters of type string and accept multiple arguments
		if len(names)==0{
			fmt.Println("Zero argument")
		}
		for _, v := range names {
			fmt.Println(v)
		}
	
}
	

//|f6|
type person struct{
	first string
}
type secretAgent struct{
	person
	ltk bool
}
func (p person) speak(){//(r reciever) take the type of the variable who call it
	fmt.Println("Hi I am ", p.first)
}
func (sa secretAgent) speak(){//(r reciever) take the type of the variable who call it
	fmt.Println("Hi I am secret agent", sa.first)
}

//Interface  |f18|
type human interface{
	speak()
}
func saySomethings(h human){
	h.speak()
}


//Stringer |f7|

type book struct {
	title string 
}

func (b book) String() string{
	return fmt.Sprint("The title of the booke is  ",b.title)
}
type count int 

func (c count) String() string{
	return fmt.Sprint("This is the number ",strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer){
	log.Println("LOG FROM 138", s.String())
}

func main(){
	//a named function with a identifier
	//func (r reciever) nameFunc(p parameters(s)) (return(s)){code}

	//anonymous function
	//func(p parameters(s)) (return(s)){code}

	

	//variadic paramaters 

	fmt.Println("---VARIADIC PARAMATERS---")

	x1:= []string{"cow","bat","frog"}
	call("Cane")
	call("dog","cat","rat")
	call()
	call(x1...)//... open a slice for multiple arguments

	fmt.Println("---END VARIADIC PARAMATERS---")

	// Methods and interface |f6| |f18|

	fmt.Println("---METHODS AND INTERFACE---")

	p:= new(person)
	p.first="Luca"
	p.speak()
	sa := secretAgent{
		person: person{
			first:"James",
		},
		ltk: true,
	}
	saySomethings(sa)

	fmt.Println("---END Methods and interface---")

	

	//Stringer and logs |f7|

	fmt.Println("---STRINGER LOGS---")

	b:= book{
		title: "Book",
	}
	var n count=42
	fmt.Println(b)
	fmt.Println(n)
	logInfo(b)
	logInfo(n)
	
	fmt.Println("---END STRINGER LOGS---")

	//anonymous function 

	fmt.Println("---ANONYMOUS FUNCTION---")

	func(){
		fmt.Println("Anonymous function runninig")
	}()
	func(s string){
		fmt.Println("My name is",s)
	}("Tod")

	fmt.Println("---END ANONYMOUS FUNCTION---")

	//return a function     |f8|

	fmt.Println("---RETURN A FUNCTION---")

	y := bar()//y assume the return of the func
	fmt.Println(y())

	fmt.Println("---END RETURN A FUNCTION---")

	//pass a func as arguments  |f9|

	fmt.Println("---PASS A FUNC AS ARGUMENTS---")

	fmt.Println(doMath(12,2,add))
	fmt.Println(doMath(12,2,sub))

	fmt.Println("---END PASS A FUNC AS ARGUMENTS---")

	//closure |f10|

	fmt.Println("---CLOSURE---")

	f:=increment()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("---END CLOSURE---")

	//recursion     |f11|

	fmt.Println("---RECURSION---")

	fmt.Println("The factorial of 4 it's",factoiral(4))

	fmt.Println("---END RECURSION---")
}
	//func return func  |f8|
	func bar() func()int{
		return func() int {
			return 42
		}
	}

//callback  |f9|
func add(a,b int)int{
	return a+b
}
func sub(a,b int)int{
	return a-b
}
func doMath(a,b int, f func(int, int) int)int{
	return f(a,b)
}

//closure    |f10|
func increment() func()int{
	x:=0
	return func() int { //the function is close in this function
		x++
		return x
	}
}
//recursion   |f11|
func factoiral(n int) int{
	if n==0{
		return 1
	}
	return n * factoiral(n-1)
}
