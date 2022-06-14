package main

import "fmt"

func main() {
	// Operadores: + - * /
	// %

	//var a int = 23
	var a float64 = 23.0
	//var b int = 22
	var b float64 = 22

	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b)
	fmt.Println("a / b = ", a/b)
	//fmt.Println("\na % b = ", a % b) // se os tipos forem inteiros
	fmt.Println("a % b = ", int(a)%int(b))
}
