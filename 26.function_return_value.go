package main

import "fmt"

func sayHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {

	callHello := sayHello("Fajar")
	fmt.Println(callHello)

	fmt.Println(sayHello("Andi"))
	fmt.Println(sayHello("Daniel"))
}
