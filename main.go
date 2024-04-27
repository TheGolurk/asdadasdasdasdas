package main

import (
	"1/suma"
	"fmt"
)

func main() {
	resultado := suma.Sumar(1, 2)

	if resultado == 3 {
		fmt.Println("correcto")
	} else {
		fmt.Println("incorrecto")
	}

}
