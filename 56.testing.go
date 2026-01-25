package main

import "fmt"

func SumAll(number ...int) int {
	total := 0

	for _, x := range number {
		total += x
	}
	return total
}

func main() {

	fmt.Println(SumAll(1, 2, 3, 4, 5))

	person := []string{
		"Fajar", "Andi", "Jono",
	}

	fmt.Println(person)
	fmt.Println(person[0])
	fmt.Println(len(person))

}
