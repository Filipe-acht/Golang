package main

import "fmt"

func main() {
	/*var idade int = 22
	var altura int = 166*/

	/*var idade, altura int = 22, 175
	var nome string = "Filipe"*/

	/*var (
		idade int = 22
		altura int = 175
		nome string = "Filipe"
	)*/

	//inferindo
	var (
		idade  = 22
		altura = 175
		nome   = "Filipe"
	)

	fmt.Println("O meu nome é ", nome, " e tenho ", idade, " anos")
	fmt.Println("Minha altura é ", altura)

	//usando a marmota

	idade2 := 22
	nome2 := "Filipe"
	altura2 := 175

	fmt.Println("\nO meu nome é ", nome2, " e tenho ", idade2, " anos")
	fmt.Println("Minha altura é ", altura2)

}
