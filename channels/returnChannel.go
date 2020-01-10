// go run returnChannel.go

package main

import "fmt"
import "time"
import "math/rand"


func main() {
	c := boring("Boring Func")
	
	for i :=0; i <5; i++ {
	   fmt.Printf("You say: %q\n", <-c) // receieve
	
	}
	
	fmt.Printf("Leaving")
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