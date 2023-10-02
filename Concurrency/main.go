package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

//waitgroup and mutex
var wg sync.WaitGroup//init a var to wait the execution of the function to make finish the main func
//var mutex sync.Mutex
//var counter int

//atomic
var counter int64

func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())//to execute the routine in the parallelism modality
}

func main() {

	//go + func start a thread
	//without the WaitGroup the func main finish before the another thread with |f12| |f13|

	wg.Add(2)//add how much routine they must run
	go foo()
	go bar()
	wg.Wait()//wait the finish of the routines
	fmt.Println("Final counter:",counter)


	/*concurrency vs parallelism 
		concurrency = indipendly executing 
		parallelism = simultaneos executiong 
	*/


}
//|f12|
func foo() {

	
	for i := 0; i < 45; i++ {
		fmt.Println("foo:",i)
		time.Sleep(time.Duration(2*time.Millisecond))
		//mutex
		/*mutex.Lock()//the variable is modificable to one routine at a time
		counter++
		fmt.Println("Foo counter: ",counter)
		mutex.Unlock()*/

		//atomic counter
		atomic.AddInt64(&counter,1)
	}
	wg.Done()//To take it off a routine who main must wait
}
//|f13|
func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:",i)
		time.Sleep(time.Duration(3*time.Millisecond))
		//mutex
		/*
		mutex.Lock()//the variable is modificable to one routine at a time
		counter++
		fmt.Println("Bar counter: ",counter)
		mutex.Unlock()*/

		//atomic counter
		atomic.AddInt64(&counter,1)
	}
	wg.Done()
}
