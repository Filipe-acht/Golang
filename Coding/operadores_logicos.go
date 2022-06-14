package main

import (
	"fmt"
)

func main() {

	var x int = 1
	var y int = 5

	// && -> and
	fmt.Println("x > 3 && x < 7:", x > 3 && x < 7)

	// || -> or
	fmt.Println("y < 3 || y < 7:", y < 3 || y < 7)

	// ! -> not
	fmt.Println("!false:", !false)
}
