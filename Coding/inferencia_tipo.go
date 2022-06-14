package main

import "fmt" // implementa  formatação para I/O

func main() {
	// var i int
	// var j = i // j vai ser tipo int
	// fmt.Printf("Tipo: %T \nValor: %v", j, j)

	//var i = 22.98 // inferir como float64
	var i = 0.64 + 0.3i // inferir como número complexo
	
	fmt.Printf("Tipo: %T  \nValor: %v", i, i)
}
