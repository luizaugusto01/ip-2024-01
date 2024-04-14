package main

import "fmt"

func main() {
    var a, b, c, d float64

    // Leitura dos elementos da matriz
    fmt.Println("Digite os quatro elementos da matriz 2x2:")
    fmt.Scanln(&a, &b, &c, &d)

    // Cálculo do determinante
    determinante := a*d - b*c

    // Impressão do determinante
    fmt.Printf("O VALOR DO DETERMINANTE E = %.2f\n", determinante)
}