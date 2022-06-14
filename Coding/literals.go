package main

import (
	"fmt"
)

func main() {
	x := "oi bom dia\n como vai?\tespero \"que\" tudo bem." //interpreted string literals
	y := `"oi bom dia 
			\tcomo vai? 
				\tespero 
				\"que sim \" "` //raw string literals
	fmt.Println(x)
	fmt.Println(y)
}
