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
	fmt.Println(sumAll(110, 20, 45, 67, 78, 77, 7))
	fmt.Println(sumAll(110, 20, 45, 67, 78, 76, 7))
	fmt.Println(sumAll(110, 20, 45, 67, 78, 77, 7))
	fmt.Println(sumAll(110, 20, 45, 67, 78, 74, 7))
	fmt.Println(sumAll(110, 20, 45, 67, 78, 75, 7))

	numbers := []int{10, 20, 30, 34, 50}
	fmt.Println(sumAll(numbers...))
}
