package main

import "fmt"

func random() interface{} {
	return 1000
}

func main() {
	tipeData := random()
	switch value := tipeData.(type) {
	case string:
		fmt.Println("Ini adalah string", value)
	case int:
		fmt.Println("Ini adalah Int", value)
	default:
		fmt.Println("Unknown", value)
	}

}
