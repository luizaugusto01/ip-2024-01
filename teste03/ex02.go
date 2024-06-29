package main

import (
    "fmt"
)

// troca troca os elementos de índice i e j no vetor.
func troca(vetor []int, i, j int) {
    vetor[i], vetor[j] = vetor[j], vetor[i]
}

// trocaOpostosSeMenor verifica e troca elementos opostos no vetor se as condições forem atendidas.
func trocaOpostosSeMenor(vetor []int, n int) {
    for i := 0; i < n/2; i++ {
        j := n - 1 - i
        if vetor[i] < vetor[j] {
            troca(vetor, i, j)
        }
    }
}

func main() {
    var t int
    fmt.Scan(&t)

    for t > 0 {
        var n int
        fmt.Scan(&n)

        vetor := make([]int, n)
        for i := 0; i < n; i++ {
            fmt.Scan(&vetor[i])
        }

        trocaOpostosSeMenor(vetor, n)

        for i, val := range vetor {
            if i > 0 {
                fmt.Print(" ")
            }
            fmt.Print(val)
        }
        fmt.Println()

        t--
    }
}
