// go run nolockstep.go

package main

import "fmt"
import "time"
import "math/rand"


func main() {
	c := fanIn2(boring("Joe"), boring("Ann"))
	
	for i :=0; i <10; i++ {
	   fmt.Printf("You say: %q\n", <-c) // receieve
	
	}
	
	fmt.Printf("Leaving")
}

func fanIn(v1, v2 <-chan string) <-chan string {
	c := make(chan string)

	go func() { for { c <- <-v1}}()
	go func() { for { c <- <-v2}}()

	return c
}

func fanIn2(v1, v2 <-chan string) <-chan string {
	c := make(chan string)

	go func() { 
		for {
			select {
			case s := <-v1: c <- s
			case s := <-v2: c <- s
			}
		}	
	}()
	
	return c
}

func boring(msg string) <-chan string { // returns recieve only channelt of strings

	c := make(chan string)

	go func(){
		for i :=0; ; i++{
			c <- fmt.Sprintf("%s %d", msg, i)
		  
			  time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		  }
	}()

	return c // returns the channel to the caller.
}