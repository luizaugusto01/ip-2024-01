package main

import "fmt"

func main() {
    var n int
    fmt.Println("Digite o número de temperaturas em Fahrenheit a serem convertidas:")
    fmt.Scanln(&n)

    for i := 0; i < n; i++ {
        var fahrenheit float64
        fmt.Printf("Digite a temperatura em Fahrenheit %d: ", i+1)
        fmt.Scanln(&fahrenheit)

        // Converte Fahrenheit para Celsius
        celsius := (fahrenheit - 32) * 5 / 9

        // Imprime a temperatura convertida
        fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", fahrenheit, celsius)
    }
}