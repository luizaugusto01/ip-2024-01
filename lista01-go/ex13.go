package main

import "fmt"

func main() {
    var nota float64

    // Leitura da nota
    fmt.Println("Digite a nota (entre 0 e 10 com uma casa decimal):")
    fmt.Scanln(&nota)

    var conceito string

    // Conversão da nota para conceito
    switch {
    case nota >= 9.0:
        conceito = "A"
    case nota >= 7.5:
        conceito = "B"
    case nota >= 6.0:
        conceito = "C"
    default:
        conceito = "D"
    }

    // Impressão da nota e do conceito correspondente
    fmt.Printf("NOTA = %.1f CONCEITO = %s\n", nota, conceito)
}