package main

import (
    "fmt"
    "strconv"
)

func main() {
    var n1, n2, n3 int
    fmt.Println("Digite os três números inteiros separados por espaço:")
    fmt.Scanln(&n1, &n2, &n3)

    // Verifica se algum dos números tem mais de um dígito
    if temMaisDeUmDigito(n1) || temMaisDeUmDigito(n2) || temMaisDeUmDigito(n3) {
        fmt.Println("DIGITO INVALIDO")
        return
    }

    // Calcula o número resultante da concatenação dos três números
    numero := n1*100 + n2*10 + n3
    // Calcula o quadrado do número
    quadrado := numero * numero

    // Imprime o número e seu quadrado
    fmt.Printf("%d, %d\n", numero, quadrado)
}

// Função para verificar se um número tem mais de um dígito
func temMaisDeUmDigito(num int) bool {
    return num < 10 || num > 99
}