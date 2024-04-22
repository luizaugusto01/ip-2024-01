package main

import "fmt"

func main() {
    var n, i, K int
    var s float64

    fmt.Scan(&n, &i, &K, &s)

    fmt.Println("Tabuada de soma:")
    for j := 0; j < K; j++ {
        resultado := float64(n) + float64(i+j)*s
        fmt.Printf("%d + %.2f = %.2f\n", n, float64(i+j)*s, resultado)
    }

    fmt.Println("Tabuada de subtracao:")
    for j := 0; j < K; j++ {
        resultado := float64(n) - float64(i+j)*s
        fmt.Printf("%d - %.2f = %.2f\n", n, float64(i+j)*s, resultado)
    }

    fmt.Println("Tabuada de multiplicacao:")
    for j := 0; j < K; j++ {
        resultado := float64(n) * float64(i+j)*s
        fmt.Printf("%d * %.2f = %.2f\n", n, float64(i+j)*s, resultado)
    }

    fmt.Println("Tabuada de divisao:")
    for j := 0; j < K; j++ {
        resultado := float64(n) / float64(i+j)*s
        fmt.Printf("%d / %.2f = %.2f\n", n, float64(i+j)*s, resultado)
    }
}