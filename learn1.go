package main

import "fmt"

func hello(a string, b string) string {
	return a + b
}

func main() {
	var name string
	fmt.Scanln(&name)
	fmt.Println(hello(name, name))
}
