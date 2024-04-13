package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    var notasString string
    fmt.Println("Digite as três notas do aluno separadas por espaço:")
    fmt.Scanln(&notasString)

    notas := strings.Split(notasString, " ")
    var total float64
    for _, nota := range notas {
        notaFloat, _ := strconv.ParseFloat(nota, 64)
        total += notaFloat
    }

    media := total / float64(len(notas))
    fmt.Printf("MEDIA = %.2f\n", media)

    if media >= 6.0 {
        fmt.Println("APROVADO")
    } else {
        fmt.Println("REPROVADO")
    }
}