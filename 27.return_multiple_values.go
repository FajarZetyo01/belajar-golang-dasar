package main

import "fmt"

func getFullname() (string, string) {
	return "Fajar", "Setyo"
}
func main() {
	firstname, lastname := getFullname()
	fmt.Println(firstname, lastname)
}
