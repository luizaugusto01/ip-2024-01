package main

import "fmt"

var numero int

func main() {

	fmt.Printf("digite um numero \n")
	fmt.Scanln(&numero)

	fatorial := make([]int, numero)

	fat := 1

	for i := 0; i < numero; i++ {

		fatorial[i] = (numero - i)

		fat *= fatorial[i]

	}

	fmt.Println(fat)

}
