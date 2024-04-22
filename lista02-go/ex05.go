package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)

    // Lê a sequência de números inteiros
    sequencia := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&sequencia[i])
    }

    // Inicializa o comprimento do segmento crescente atual e o máximo
    comprimentoAtual := 1
    comprimentoMaximo := 1

    // Percorre a sequência para encontrar o comprimento máximo do segmento crescente
    for i := 1; i < n; i++ {
        if sequencia[i] > sequencia[i-1] {
            comprimentoAtual++
        } else {
            // Se o segmento terminar, atualiza o comprimento máximo, se necessário
            if comprimentoAtual > comprimentoMaximo {
                comprimentoMaximo = comprimentoAtual
            }
            comprimentoAtual = 1
        }
    }

    // Verifica se o comprimento atual é maior que o máximo
    if comprimentoAtual > comprimentoMaximo {
        comprimentoMaximo = comprimentoAtual
    }

    // O comprimento do segmento é o tamanho do maior segmento menos um
    resultado := comprimentoMaximo - 1

    fmt.Printf("O comprimento do segmento crescente maximo e: %d\n", resultado)
}