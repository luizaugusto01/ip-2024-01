package main

import (
	"fmt"
)

var n1, n2, n3, numero int

func main() {

	fmt.Println("digite os números a, b, c ")
	fmt.Scan(&n1, &n2, &n3)

	if n1 < 0 || n1 > 9 || n2 < 0 || n2 > 9 || n3 < 0 && n3 > 9 {
		fmt.Println("digito invalido")
		return
	}

	numero = (n1*100 + n2*10 + n3)
	fmt.Printf("%v, %v", numero, numero*numero)

}
