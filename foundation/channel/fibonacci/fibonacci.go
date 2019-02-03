package main

import "fmt"

func main() {
	c := make(chan int, 5)
	go fibonacci(4, c)
	for i := 0; i < 4; i++ {
		fmt.Println(<- c)
	}
}

func fibonacci(num int, c chan int)  {
	x, y := 1, 1
	for i := 0; i < num; i++ {
		c <- y
		x, y = y, x + y
	}
	close(c)
}
