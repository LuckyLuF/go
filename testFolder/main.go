package main

import "fmt"

func main() {
	//test    |f19|
	fmt.Println("---FIRST TEST---")
	fmt.Println("2+3=",sum(2,3))
	fmt.Println("4+7=",sum(4,7))
	fmt.Println("5+9=",sum(5,9))

	fmt.Println("---END FIRST TEST---")

	//test    |f20|
	fmt.Println("---SECOND TEST---")
	fmt.Println("2+3=",sum(2,3))
	fmt.Println("4+7=",sum(4,7))
	fmt.Println("5+9=",sum(5,9))

	fmt.Println("---END SECOND TEST---")
}


//|f19|
func sum(x1 ...int) int {	
	sum := 0
	for _, v := range x1 {
		sum += v
	}
	return sum
	
}

//|f20|
func mySum(x1 ...int) int {	
	sum := 0
	for _, v := range x1 {
		sum += v
	}
	return sum
	
}