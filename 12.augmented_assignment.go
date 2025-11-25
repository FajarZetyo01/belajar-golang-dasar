// EXAMPLE

// a = a + 10 ->>>>>> a+=10
// a = a / 10 ->>>>>> a/=10
// a = a * 10 ->>>>>> a*=10

package main

import "fmt"

func main() {
	a := 20
	a += 10

	fmt.Println(a)
	a *= 10
	fmt.Println(a)

	//UNARY OPERATOR

	var j = 15
	j++ // J = 15 + 1
	fmt.Println(j)
	j-- // J = 15 - 1
	fmt.Println(j)
}
