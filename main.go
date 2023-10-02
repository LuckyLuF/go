package main

import (
	//"fmt"
)

//"math/rand"
//"time"

func main(){
	/*//Channels, defer, routines, select with |f1| and |f2|
	
	ch1:= make(chan int)
	ch2:= make(chan int)
	defer fmt.Println(ch1)
	fmt.Printf("ch1: %d\n", ch1)
	var a = 32
	const cavallo string = "Cavallo"
	fmt.Println(cavallo)
	call(a)
	for i:=0; i<10; i++{
		if i%2!=0{
			continue
		}
		fmt.Println("counting even numbers: ",i)
	}
	d1:= time.Duration(rand.Int63n(2))
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch2 <- 47
	}()
	select{
		case vi:= <-ch1:
		fmt.Println("1",vi)
		case v2:= <- ch2:
			fmt.Println("2",v2)
	}
	if x:=Intn(250); x>0 && x<100{
		fmt.Println("Woooooooooow")}*/
	
	/*//Pointer
	
	i:=10
	p:= &i
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(i)
	i=20
	x:=0xc00000a0b8
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(i)
	fmt.Println(x)*/
	/*//Array, cap and len
	
	a:= [3]int{}
	x:= a[1:3]
	b:= [...]string{"Hey", "Uajiu","sono", "Golang"}//[...] capisce quanti elementi sono
	// b[10]="Ciao" //non funziona perché out of bounds
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))*/

	/*//Slice, append, delete with |1|, make 
	
	xs:=[]int {0,1,2,3,4,5,6,7,8,9,10,11,12}
	fmt.Printf("%v",xs)
	xs = append(xs, 100)
	fmt.Println(xs)

	fmt.Println(xs[4:5])
	xs = append(xs[:4], xs[5:]...)  //[:4] fino al numero 4 ma non incluso quindi 3, [5:] dal numero 5 incluso   |1|
	fmt.Printf("xs: %v\n", xs)

	xw:= make([]int, 0,2) // []int slice int, 0 lunghezza iniziale, 2 capacità iniziale (che può aumentare per multipli dell'inizializzazione)
	xw = append(xw, 1,2,3,4,5,6,7,8)
	fmt.Println(xw)*/


	/*//{}Block the scope

		var x int
		fmt.Println(x)
		} //block the scope between the brackets it ins't accesible from outside
		fmt.Println(x) //don't found*/

	//[inclusive:esclusive]

	/*for _, v := range xs { //don't use the output use _ |||   for i,v:= range xs
		fmt.Println(v)
	}*/

	//slice point an array

	/*//copy slice and the slice point to another slice

		a:=[]int{1,2,3,4,5}
		b:= a //a->array, b->a b->array(a)
		//modify b, modify a
		c:= make([]int, len(a), 10)
		d:= []int{}
		a[0]=7
		b[1]=9
		fmt.Println(a)
		fmt.Println(b)
		copy(d, a)
		copy(c, a)
		c[2]=10
		fmt.Println(c)
		fmt.Println(len(c))
		fmt.Println(cap(c))
		c = append(c, 1,2)
		fmt.Println(c)
		fmt.Println(d)
		fmt.Println(len(d))
		fmt.Println(cap(d))
		println("a:",a)
		println("b:",b)
		println("c:",c)println("d:",d)*/

	/*//Pass by value and by reference with |f3| and |f4| function

		x:=2
		fmt.Println("Init x:",x)
		ByValue(x)
		fmt.Println("After ByValue function x:",x)
		ByReference(&x)
		fmt.Println("After ByReference function x:",x)*/

	/*//double slices
	
	jb:=[]string{"Cavallo","Orso","Pappagallo"}
	jm:=[]string{"Galoppa","Roar","Copia"}
	//ji:=[]int{1,2,3}
	
	xq:=[][]string{jb,jm}

	fmt.Println(xq)*/

	/*//map create and delete, switch, ok comma
	m:= map[string]int{
		"Carlo":11,
		"Io":92,
		"Cavallo":32,
		"P":90,
	}
	for k, v := range m {
		switch k{
			case "Io":
				fmt.Printf("%s have %d years old\n",k,v)
				break
			case "Carlo":
				fmt.Printf("%s has %d years old\n",k,v)
			case "Cavallo":
				fmt.Printf("%s has %d years old\n",k,v)
			default:
				fmt.Printf("%s have %d years old\n",k,v)
		}
		fmt.Printf("%s have %d years old\n",k,v)
	}

	//or
	ma := make(map[string]int)
	ma["Cavallo"]=12
	//delete
	delete(m, "Carlo")
	fmt.Println("After delete Carlo")
	for k, v := range m {
		fmt.Printf("%s have %d years old\n",k,v)
	}

	mw:= make(map[string]int)
	q, ok:=mw["tudnd"]
	fmt.Println(q,ok)//ok false because the key don't exist
	fmt.Println(mw["tunde"])//don't exist return 0
	mw["tunde"]++ //create the key tunde and increment the value 0 -> 1
	fmt.Println(mw["tunde"])
	for k, v := range mw {
		fmt.Println(k,v)
	}*/

	/*//create a map wih a slice in the value
	m:= make(map[string][]int)
	m["Dog"] = []int{1,2,3,4,5}
	m["Cat"] = []int{6,7,8,9,10}
	m["Bee"] = []int{11,12,13,14,15}

	for k, v := range m {//like array
		fmt.Println("The key is",k,", the values are:",v)
	}
	for k, v := range m {//detail
		fmt.Println("The key",k,"have the values")
		for value := range v{
			fmt.Println(value)
		}
	}
	delete(m, "Bee")
	fmt.Println("After delete Bee")
	fmt.Println(m["Dog"][2])
	m["Dog"][2]=200
	m["Rat"] = []int{15,16,17,18,19,20}
	for k, v := range m{//detail
		fmt.Println("The key",k,"have the values")
		fmt.Println(v)
		for _,value := range v{
			fmt.Println(value)
		}
	}*/

	/*//Struct

	type person struct{
		first string 
		last string
		age int
	}
	type agent struct{
		person person
		working bool
	}
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
	*/

	/*//variadic paramaters 
	x1:= []string{"cow","bat","frog"}
	call("Cane")
	call("dog","cat","rat")
	call()
	call(x1...)//... open a slice for multiple arguments
	*/

	/*// Methods |f6|
	p:= new(person)
	p.first="Luca"
	p.speak()
	*/


	

}
//func (r reciever) nameFunc(p parameters(s)) (return(s)){code}

/*//|f6|
type person struct{
	first string
}
func (p person) speak(){//(r reciever) take the type of the variable who call it
	fmt.Println("Hi I am ", p.first)
}*/

//|f5|
/*func call(names ...string){ // ... -> the function have multiple parameters of type string and accept multiple arguments
	if len(names)==0{
		fmt.Println("Zero argument")
	}
	for _, v := range names {
		fmt.Println(v)
	}

}*/


/*func pass by value and reference |f3| |f4|

//|f3|
func ByValue(x int){//assume the value
	x=x+5
	fmt.Println("Inside ByValue function x:",x)
}
//|f4|
func ByReference(x *int){ //Point to value
	*x=*x+5
	fmt.Println("Inside ByReference function x:",*x)}*/

/*//|f1| |f2|
//|f1|
func call(a int){
	fmt.Println("Calling",a,"People")

}
//|f2|
func Intn(n int) int{
	return rand.Intn(n)}*/
