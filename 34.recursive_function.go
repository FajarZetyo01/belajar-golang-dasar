package main

import "fmt"

func recursiveFactorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursiveFactorial(value-1)
	}
}

func main() {

	value := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	fmt.Println(value)
	fmt.Println(recursiveFactorial(10))

}
