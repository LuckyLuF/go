package main

import "fmt"


func addI(a,b int) int{
	return a+b
}
func addF(a,b float64) float64{
	return a+b	
}
func addT[T int | float64] (a,b T) T{
	return a+b
}
type myNumber interface{
	~int | float64  //~ accept all value who have a int under itself
}
func addTi[T myNumber] (a,b T) T{
	return a+b
}
type myAllias int
func main(){

	var n myAllias=42
	fmt.Println(addI(1,2))
	fmt.Println(addF(1.12,2.32))
	
	fmt.Println(addT(1,2))
	fmt.Println(addT(1.12,2.32))

	fmt.Println(addTi(1,2))
	fmt.Println(addTi(1.12,2.32))

	fmt.Println(addTi(n,2))  //need the ~ to accept myAllias 
	fmt.Println(addTi(1.12,2.32))

	//addT[int]() // int is in a possible type
	//addT[float64]() // float64 is a the possible type
	//addT[string]() // string isn't a the possible type

	//concrete vs interface type
}