package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func main() {
	array := []int{1, 2, 3, 45, 453}
	c := make(chan int)
	go sum(array, c)
	go sum(array[:len(array)/2], c)

	x, y := <-c, <-c
	fmt.Println(x)
	fmt.Println(y)
}
