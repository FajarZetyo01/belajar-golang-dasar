package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Youre blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {

	//CARA 1
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Eko", blacklist)
	registerUser("Anjing", blacklist)

	//CARA 2
	registerUser("Fajar", func(name string) bool {
		return name == "Anjing"
	})
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})

}
