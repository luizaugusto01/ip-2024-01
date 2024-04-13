package main

import (
    "fmt"
)

func main() {
    var casosTeste int
    fmt.Scanln(&casosTeste)

    for i := 1; i <= casosTeste; i++ {
        var publicoTotal, percentPopular, percentGeral, percentArquibancada, percentCadeiras float64
        fmt.Scanln(&publicoTotal, &percentPopular, &percentGeral, &percentArquibancada, &percentCadeiras)

        rendaTotal := (percentPopular * 0.01 * 1.0) +
                      (percentGeral * 0.01 * 5.0) +
                      (percentArquibancada * 0.01 * 10.0) +
                      (percentCadeiras * 0.01 * 20.0)

        fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i, rendaTotal)
    }
}