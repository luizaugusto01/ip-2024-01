package main

import "fmt"

func main() {
    var raio, altura float64

    // Leitura do raio da lata
    fmt.Println("Digite o raio da lata (em metros):")
    fmt.Scanln(&raio)

    // Leitura da altura da lata
    fmt.Println("Digite a altura da lata (em metros):")
    fmt.Scanln(&altura)

    // Cálculo da área da superfície da lata
    area := 2 * 3.14 * raio * (raio + altura)

    // Custo do alumínio por m²
    custoAluminioPorMetroQuadrado := 100.00

    // Cálculo do custo da lata
    custoLata := area * custoAluminioPorMetroQuadrado

    // Impressão do custo da lata
    fmt.Printf("O VALOR DO CUSTO E = %.2f\n", custoLata)
}