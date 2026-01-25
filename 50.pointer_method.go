package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

//KETIKA DIFUNCTION ADA METHOD LEBIH DISARANKAN MENGGUNAKAN POINTER pada type Struct

func main() {
	fajar := Man{"Fajar"}
	fajar.Married()
	fmt.Println(fajar.Name)

}
