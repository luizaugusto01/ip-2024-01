package main

import (
    "fmt"
)

func calcularNotaFinal(provas []float64, listas []float64, trabalhoFinal float64) float64 {
    var somaProvas, somaListas float64

    for _, nota := range provas {
        somaProvas += nota
    }

    for _, nota := range listas {
        somaListas += nota
    }

    mediaProvas := somaProvas / float64(len(provas))
    mediaListas := somaListas / float64(len(listas))

    notaFinal := 0.7*mediaProvas + 0.15*mediaListas + 0.15*trabalhoFinal

    return notaFinal
}

func main() {
    var matricula int
    for {
        fmt.Scan(&matricula)
        if matricula == -1 {
            break
        }

        var provas [8]float64
        var listas [5]float64
        var trabalhoFinal, presenca float64

        for i := 0; i < 8; i++ {
            fmt.Scan(&provas[i])
        }

        for i := 0; i < 5; i++ {
            fmt.Scan(&listas[i])
        }

        fmt.Scan(&trabalhoFinal, &presenca)

        notaFinal := calcularNotaFinal(provas[:], listas[:], trabalhoFinal)

        var situacaoFinal string
        if notaFinal >= 6 && (presenca/128)*100 >= 75 {
            situacaoFinal = "APROVADO"
        } else if notaFinal < 6 && (presenca/128)*100 >= 75 {
            situacaoFinal = "REPROVADO POR NOTA"
        } else if notaFinal >= 6 && (presenca/128)*100 < 75 {
            situacaoFinal = "REPROVADO POR FREQUENCIA"
        } else {
            situacaoFinal = "REPROVADO POR NOTA E POR FREQUENCIA"
        }

        fmt.Printf("Matricula: %d, Nota Final: %.2f, Situacao Final: %s\n", matricula, notaFinal, situacaoFinal)
    }
}