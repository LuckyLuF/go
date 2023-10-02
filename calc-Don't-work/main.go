package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"strconv"

	//"golang.org/x/tools/go/analysis/passes/printf"
)

func main() {              //Don't work
	var result float64
	var first float64
	var second float64
	var operator string
	fmt.Println("Inserire il primo numero")
	//reader:=bufio.NewReader(os.Stdin)
	// first, _:= reader.ReadString('\n')
	// nfirst, _  :=strconv.ParseFloat(first, 64)
	fmt.Scan(&first)
	fmt.Println("Inserire il secondo numero")
	fmt.Scan(&second)
	//second, _:= reader.ReadString('\n')
	//nsecond, _  :=strconv.ParseFloat(second, 64)
	fmt.Println("Inserire l'opratore tra *(moltiplicazione), /(divisione),-(sottrazione), +(addizione)")
	fmt.Scan(&operator)
	//reader.ReadString('\n')
	//operator, _:= reader.ReadString('\n')
	for result!=0{
		switch operator{
		case "*\n":
			result = first * second
		case "+\n":
			result = first + second
		case "-\n":
			result = first - second
		case "/\n":
			result = first / second
		default:
			fmt.Println("Inserire uno di questi operatori: *  /  -  +")
		}
	}
	// fmt.Println(first,nfirst,second,nsecond,operator)
	// fmt.Printf("%T%T%T%T%T",first,nfirst,second,nsecond,operator)
	// fmt.Println("Il risultato dell'operazione",nfirst,operator,nsecond,"è",result)
	fmt.Println(first,second,operator,result)
}