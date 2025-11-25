package main

import "fmt"

func panggilBiodata(firstName string, middleName string, lastName string, umur int) {
	fmt.Println("Biodata kamu : ", firstName, middleName, lastName, umur)
}

func main() {

	panggilBiodata("Fajar", "Setyo", "Pambudi", 25)
	panggilBiodata("Andi", "Setyo", "Pambudi", 30)
}
