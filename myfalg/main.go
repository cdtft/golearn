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
	flag.Parse()
	fmt.Printf("hello %s /n", name)
	internal.Hello()
}
