package main

import "fmt"

var name string = "Daan"

func main() {
	fmt.Println(hello(name))
}

func hello(name string) string {
	return "hello " + name
}
