package main

import (
	"fmt"
)

var casosteste, pessoatotal int
var percpopular, pergeral, perarquibancada, percadeiras, rendatotal float64

func main() {

	fmt.Print("Digite o número de casos de teste: ")
	fmt.Scan(&casosteste)

	for i := 1; i <= casosteste; i++ {
		fmt.Println("digite o número de pessoas: ")
		fmt.Scan(&pessoatotal)
		fmt.Println("digite percpopular: ")
		fmt.Scan(&percpopular)
		fmt.Println("digite o percgeral: ")
		fmt.Scan(&pergeral)
		fmt.Println("digite o percarquibancada: ")
		fmt.Scan(&perarquibancada)
		fmt.Println("digite o número de perccadeiras: ")
		fmt.Scan(&percadeiras)

		rendatotal = (percpopular * float64(pessoatotal) * 1) + (pergeral * float64(pessoatotal) * 5) + (perarquibancada * float64(pessoatotal) * 10) + (percadeiras * float64(pessoatotal) * 20)

		fmt.Printf("A renda do jogo %v é %v\n\n", i, rendatotal)
	}
}
