package main

import "fmt"

func getFullname() (firstname, middlename, lastname string) {
	firstname = "Fajar"
	middlename = "Setyo"
	lastname = "Pambudi"

	return firstname, middlename, lastname
}

func main() {

	a, b, c := getFullname()
	fmt.Println(a, b, c)

}
