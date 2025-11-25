package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, x := range numbers {
		total += x
	}
	return total
}

func main() {
	result := sumAll(10, 20, 30, 40, 50)
	fmt.Println(result)

}
