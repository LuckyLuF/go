package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}
//ByAge implements the sort interface with Len,Swap and Less
type ByAge []Person

func (a ByAge) Len()int {return len(a)}
func (a ByAge) Swap(i,j int){a[i], a[j] = a[j], a[i]}
func (a ByAge) Less(i, j int)bool {return a[i].Age < a[j].Age} //less to more
func main() {
	//sort
	numbers := []int{6, 4, 8, 2, 0, 3, 5}
	fmt.Println("Before sort", numbers)
	sort.Ints(numbers)
	fmt.Println("After sort", numbers)

	//custom sort ByAge
	people := []Person{
		{"Henry",31},
		{"Bob",12},
		{"Carlo",34},
		{"Max",87},
	}
	fmt.Println("Before sort", people)
	sort.Sort(ByAge(people))//need the sort interface 
	fmt.Println("After sort", people)

}
