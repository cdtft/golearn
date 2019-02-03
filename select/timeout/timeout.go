package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	timeout := make(chan int)

	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				timeout <- 0
				break
			}
		}
	}()

	for {
		c <- 0
	}
	<-timeout
	fmt.Println("timeout...5!")
}
