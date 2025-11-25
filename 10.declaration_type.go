package main

import "fmt"

func main() {
	type KTPstring string

	noKTP := KTPstring("1234567890")
	fmt.Println(noKTP)

	noKTPfajar := "98674574564564"
	noKTPfajarstring := KTPstring(noKTPfajar)
	fmt.Println(noKTPfajarstring)
}
