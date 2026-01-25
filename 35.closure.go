package main

import "fmt"

func counter() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	counter1 := counter()
	counter1()
	counter1()
	counter1()
	fmt.Println(counter1())

}
