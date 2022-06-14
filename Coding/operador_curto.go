package main

import "fmt"

func main() {

	//Operador curto(declaração) atribui automaticamente o tipo com base no valor da variável
	x := 10
	y := "hallo welt"

	//fmt.Println(x, y)

	//print formatado
	fmt.Printf("x: %v, %T\n", x, x) // %v - valor, %T- tipo
	fmt.Printf("y: %v, %T", y, y)
}
