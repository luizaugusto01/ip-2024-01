package main

import "fmt"

func main() {
    var numero int

    // Leitura do número inteiro
    fmt.Println("Digite um número inteiro:")
    fmt.Scanln(&numero)

    // Verifica se o número é divisível por três e por cinco
    if numero%3 == 0 && numero%5 == 0 {
        fmt.Println("O NUMERO E DIVISIVEL")
    } else {
        fmt.Println("O NUMERO NAO E DIVISIVEL")
    }
}