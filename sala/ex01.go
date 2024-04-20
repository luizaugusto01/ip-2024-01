package main

import (
	"fmt"
)

func main() {

	var notas [3]float64
	fmt.Println("Digite as notas: ")

	for i := 0; i < 3; i++ {
		fmt.Scan(&notas[i])
	}

	for i := 0; i < 3; i++ {
		fmt.Println(notas[i])
	}

	media := (notas[0] + notas[1] + notas[2]) / 3

	fmt.Printf("%.2f \n", media)

}