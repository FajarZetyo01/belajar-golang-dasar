package main

import "fmt"

// CARA 1
func filterName(name string, filter func(string) string) {
	filterNama := filter(name)
	fmt.Println("Hello ", filterNama)
}

// CARA 2
type Filter func(string) string

//  ||Nama func|| tipe data|| returnnya string

func filterName1(name string, filter Filter) {
	filterNama1 := filter(name)
	fmt.Println("Hello ", filterNama1)
}

func filterSpam(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	//CARA KE 1
	filterName("Fajar", filterSpam)
	filter := filterSpam
	filterName("Anjing", filter)

	//CARA KE 2
	filterName1("Fajar", filterSpam)
	filter1 := filterSpam
	filterName("Daniel", filter1)
}
