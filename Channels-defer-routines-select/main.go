package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//|f1| |f2|
//|f1|
func call(a int){
	fmt.Println("Calling",a,"People")
}
//|f2|
func Intn(n int) int{
	return rand.Intn(n)
}


//|f16|
func populate(c chan int ){
	for i := 0; i < 100; i++ {
		c<-i
	}
	close(c)
}

//|f16|
func fanOutIn(c1,c2 chan int){
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
			go  func (v2 int)  {//|f17|
				c2 <- timeConsumingWork(v2)
				wg.Done()
			}(v)
	}
	wg.Wait()
	close(c2)
}

//|f17|
func timeConsumingWork(n int) int{
	time.Sleep(time.Microsecond*time.Duration(rand.Intn(500)))
	return n+rand.Intn(1000)
}


//|f14|
func send(even, odd chan<- int ){
	for i := 0; i < 100; i++ {
		if i%2==0 {
			even <- i
		} else{
			odd <- i
		}
	}
	close(even)
	close(odd)
}
//|f15|
func receive(even, odd <- chan int, fanin chan<-int){
	var wg sync.WaitGroup
	wg.Add(2)

	go func ()  {
		for v := range even {
			fanin <- v
			time.Sleep(time.Duration(3* time.Millisecond))//it'll wait because conncet at the same channels
		}
		wg.Done()
	}()
	go func ()  {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)
}


//|f12|  |f13|

//|f12|
func data()chan int{
	c:=make(chan int)
	go func ()  {
		for i := 0; i < 10; i++ {
			c<-i
		}
		close(c)
	}()
	return c
}
//|f13|
func sum(c chan int) chan int {
	var sum int
	out:= make(chan int )
	go func ()  {
		for e := range c {
			sum +=e
		}
		out<- sum
		out<- sum
		close(out)
		//out<-sum //don't work channel is closed
	}()
	return out 
}


func main() {

	//tests

	/*cha:=make(chan chan int)
	ce:=make(chan int)
	ce<- 12
	cha<-ce
	close(cha)
	close(ce)
	fmt.Println(ce)
	fmt.Println(cha)*/



	/*
	chan T          // can be used to send and receive values of type T
	chan<- float64  // can only be used to send float64s
	<-chan int      // can only be used to receive ints

	chan<- chan int    // same as chan<- (chan int)
	chan<- <-chan int  // same as chan<- (<-chan int)
	<-chan <-chan int  // same as <-chan (<-chan int)
	chan (<-chan int)
	
	*/


	//Channels, defer, routines, select with |f1| and |f2|

	fmt.Println("---CHANNELS, DEFER, ROUTINES, SELECT---")

	ch1 := make(chan int)
	ch2 := make(chan int)
	defer fmt.Println(ch1)
	fmt.Printf("ch1: %d\n", ch1)
	var a = 32
	const cavallo string = "Cavallo"
	fmt.Println(cavallo)
	call(a)
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("counting even numbers: ", i)
	}
	d1 := time.Duration(rand.Int63n(2))
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch2 <- 47
	}()
	select {
	case vi := <-ch1:
		fmt.Println("1", vi)
	case v2 := <-ch2:
		fmt.Println("2", v2)
	}
	if x := Intn(250); x > 0 && x < 100 {
		fmt.Println("Woooooooooow")
	}

	fmt.Println("---END CHANNELS, DEFER, ROUTINES, SELECT---")


	//channels

	fmt.Println("---CHANNELS---")
	c:= make(chan int)
	go func ()  {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)// remember to close the channel 
	}()

	for n:= range c{
		fmt.Println(n)
	}
	fmt.Println("---END CHANNELS---")
	//channels in the functions with |f12| |f13|
	fmt.Println("---CHANNELS IN THE FUNCTIONS---")
	q := data()
	qSum:= sum(q)
	/*for e := range qSum {
		fmt.Println(e)
	}*/
	fmt.Println(<-qSum)
	fmt.Println(<-qSum)
	fmt.Println(<-qSum)//0 the channel is void 

	fmt.Println("---END CHANNELS IN THE FUNCTIONS---")

	//fan in multiple channel write in a channel(fan in) |f14| |f15|
	fmt.Println("---FAN IN---")

	even:= make(chan int)
	odd:= make(chan int)
	fanIn:= make(chan int)

	go send(even, odd)
	go receive(even,odd,fanIn)

	for v := range fanIn{
		fmt.Println(v)
	}
	fmt.Println("about the exit")

	fmt.Println("---END FAN IN---")


	//fan in multiple channel write in a channel(fan out) |f16| |f17|
	

	fmt.Println("---FAN OUT---")

	c4:=make(chan int)
	c5:=make(chan int)

	go populate(c4)

	go fanOutIn(c4,c5)

	for v := range c5 {
		fmt.Println(v)
	}
	fmt.Println("about the exit")

	fmt.Println("---END FAN OUT---")
	
}

