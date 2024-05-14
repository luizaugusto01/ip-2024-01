package main
import "fmt"

func main() {
    var n, i, k, s float64

    fmt.Println("Digite um número n (0 a 9)")
    fmt.Scanln(&n)


    fmt.Println("digite o valor inicial i")
    fmt.Scanln(&i)

    fmt.Println("digite o número de valores k:")
    fmt.Scanln(&K)

    fmt.Println("digite o valor k:")
    fmt.Scanln(&s)

    values := make([]float64, int(k))
    for j := range values {
        values[j] = i + flaot64(j)*s
    }

    fmt.Println("tabuada de soma :")
    for _, B := n + B {
    fmt.Printf("%.2f + %.2f = %.2f\n", n, B,)
    }
}