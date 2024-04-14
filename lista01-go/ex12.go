package main

import "fmt"

func main() {
    var horas int

    // Leitura do número de horas de uso da charrete
    fmt.Println("Digite a quantidade de horas de uso da charrete:")
    fmt.Scanln(&horas)

    // Cálculo do valor a pagar
    var valorAPagar float64
    if horas >= 3 {
        valorAPagar = float64(horas / 3 * 10)
    } else {
        valorAPagar = float64(horas * 5)
    }

    // Impressão do valor a pagar
    fmt.Printf("O VALOR A PAGAR E = %.2f\n", valorAPagar)
}