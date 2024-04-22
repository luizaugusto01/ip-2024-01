package main

import "fmt"

func fatorial(n int) int {
    if n == 0 {
        return 1
    }
    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    return result
}

func main() {
    var n int
    fmt.Scan(&n)

    resultado := fatorial(n)

    fmt.Printf("%d! = %d\n", n, resultado)
}