package main

import "fmt"

func main() {
	var a int = 46
	//var b float64 = 6.4

	//var c int = int(b) // coerção
	var c float64 = float64(a)

	fmt.Println("O valor de c é ", c)
}
