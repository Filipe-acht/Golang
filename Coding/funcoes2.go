package main

import (
	"fmt"
)

// função que retorna o quadrado e cubo de um numero inteiro
func calcular(a int) (quadrado int, cubo int) { // valores de retorno (quadado int, cubo int) podem ser nomeados(só para funções curtas).
	/*var*/ quadrado /*int*/ = a * a // remover o var e o tipo
	/*var*/ cubo /*int*/ = a * a * a

	return quadrado, cubo
}

func main() {
	fmt.Println(calcular(2))
}
