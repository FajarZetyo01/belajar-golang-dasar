package main

import "fmt"

func sayGoodBye(name string) string {
	return "Hello " + name
}

func main() {
	sayGoodbye := sayGoodBye
	sayGoodbye1 := sayGoodBye
	fmt.Println(sayGoodbye("Fajar"))
	fmt.Println(sayGoodbye1("Daniel"))

}
