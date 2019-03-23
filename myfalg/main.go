package main

import (
	"demo/myfalg/internal"
	"flag"
	"fmt"
)

func init() {
	flag.StringVar(&name, "name", "everyone", "the greeting object.")
}

var name string

func main() {
	//map的空指针测试
	var myMap map[string]int

	delete(myMap, "1")
	myMap["s"] = 1

	flag.Parse()
	fmt.Printf("hello %s \n", name)
	internal.Hello()
}
