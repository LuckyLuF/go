package main

import "fmt"

func main(){

	//make(map[T]T)  or
	/* m:=map[T]T{
		key:value
	}*/
	//map create and delete, switch, ok comma
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
	}

	//create a map wih a slice in the value
	ms:= make(map[string][]int)
	ms["Dog"] = []int{1,2,3,4,5}
	ms["Cat"] = []int{6,7,8,9,10}
	ms["Bee"] = []int{11,12,13,14,15}

	for k, v := range ms {//like array
		fmt.Println("The key is",k,", the values are:",v)
	}
	for k, v := range ms {//detail
		fmt.Println("The key",k,"have the values")
		for value := range v{
			fmt.Println(value)
		}
	}
	delete(ms, "Bee")
	fmt.Println("After delete Bee")
	fmt.Println(ms["Dog"][2])
	ms["Dog"][2]=200
	ms["Rat"] = []int{15,16,17,18,19,20}
	for k, v := range ms{//detail
		fmt.Println("The key",k,"have the values")
		fmt.Println(v)
		for _,value := range v{
			fmt.Println(value)
		}
	}
}