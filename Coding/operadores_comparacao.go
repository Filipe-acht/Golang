package main

import (
	"fmt"
)

func main() {
	// > < >= <=
	// == !=
	// retorno = booelano

	var a int = 23
	var b int = 7

	fmt.Println("a >  b: ", a > b)
	fmt.Println("a <  b: ", a < b)
	fmt.Println("a >= b: ", a >= b)
	fmt.Println("a <= b: ", a <= b)

	fmt.Println("3 == 4:", 3 == 4) //igual a
	fmt.Println("21 == 21:", 21 == 21)

	fmt.Println("a != b:", a != b) // diferente de
}
