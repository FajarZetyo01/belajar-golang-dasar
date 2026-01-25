package main

import "fmt"

type Address struct {
	Name, Address, Province string
}

func main() {

	Address1 := new(Address)
	Address2 := Address1

	Address2.Name = "Daniel Prasetyo"

	fmt.Println(Address1)
	fmt.Println(Address2)

}
