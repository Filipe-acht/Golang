package main


import "fmt"

func main(){
	var altura = 1.75  //go vai inferenciar como float
	var aberto bool = true

	fmt.Printf("Tipo: %T \nValor: %v", altura, altura)
	fmt.Printf("\nTipo: %T \nValor: %v", aberto, aberto)

}