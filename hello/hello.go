package main

import (
	"fmt"
)

//HelloWorld func
func HelloWorld(name string) string {
	var identifier string
	if name != "" {
		identifier = name
	} else {
		identifier = "World"
	}
	return "Hello, " + identifier
}

func main() {
	fmt.Println(HelloWorld(""))
}
