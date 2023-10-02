package main

import "fmt"

func main(){

	//make([]T, length, capacity) or s:=[]int{}
	//Slice, append, delete with |1|, make 
	
	xs:=[]int {0,1,2,3,4,5,6,7,8,9,10,11,12}
	fmt.Printf("%v",xs)
	xs = append(xs, 100)
	fmt.Println(xs)

	fmt.Println(xs[4:5])
	xs = append(xs[:4], xs[5:]...)  //[:4] fino al numero 4 ma non incluso quindi 3, [5:] dal numero 5 incluso   |1|
	fmt.Printf("xs: %v\n", xs)

	xw:= make([]int, 0,2) // []int slice int, 0 lunghezza iniziale, 2 capacità iniziale (che può aumentare per multipli dell'inizializzazione)
	xw = append(xw, 1,2,3,4,5,6,7,8)
	fmt.Println(xw)


	//{}Block the scope
		{
		var x int
		fmt.Println(x)
		} //block the scope between the brackets it ins't accesible from outside
		//fmt.Println(x) //don't found

	//[inclusive:esclusive]

	for _, v := range xs { //don't use the output use _ |||   for i,v:= range xs
		fmt.Println(v)
	}

	//slice point an array

	//copy slice and the slice point to another slice

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
		println("c:",c)
		println("d:",d)

		//double slices
	
		jb:=[]string{"Cavallo","Orso","Pappagallo"}
		jm:=[]string{"Galoppa","Roar","Copia"}
		//ji:=[]int{1,2,3}
	
		xq:=[][]string{jb,jm}

		fmt.Println(xq)
}