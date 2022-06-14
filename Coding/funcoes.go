package main

import "fmt"

// usa-se func para criar uma função, () - lista de parametros
//func soma(x, y int) int -> significa que todas são do tipo int

func soma(x int, y int) int { //int é o retorno da função
	return x + y
}

func main() {
	//soma(3, 4) //chamand a função: 3 -> x e 4 -> y

	fmt.Printf("Valor da função: %v",soma(3,4))
}
