package main

import "fmt"

func main() {
    var contaCliente int
    var consumoAgua float64
    var tipoConsumidor rune

    // Leitura da entrada: conta do cliente, consumo de água e tipo de consumidor
    fmt.Println("Digite a conta do cliente, o consumo de água (em metros cúbicos) e o tipo de consumidor (R - Residencial, C - Comercial, I - Industrial), separados por espaço:")
    fmt.Scanln(&contaCliente, &consumoAgua, &tipoConsumidor)

    // Calcula o valor da conta com base no tipo de consumidor
    var valorConta float64
    switch tipoConsumidor {
    case 'R':
        valorConta = 5.0 + (0.05 * consumoAgua)
    case 'C':
        if consumoAgua <= 80 {
            valorConta = 500.0
        } else {
            valorConta = 500.0 + (0.25 * (consumoAgua - 80))
        }
    case 'I':
        if consumoAgua <= 100 {
            valorConta = 800.0
        } else {
            valorConta = 800.0 + (0.04 * (consumoAgua - 100))
        }
    }

    // Imprime a conta do cliente e o valor a ser pago
    fmt.Printf("CONTA = %d\n", contaCliente)
    fmt.Printf("VALOR DA CONTA = %.2f\n", valorConta)
}