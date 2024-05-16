package main

import "fmt"

func main() {
	var n, i, k, s float64

	fmt.Println("Digite um número n (0 a 9)")
	fmt.Scanln(&n)

	fmt.Println("digite o valor inicial i")
	fmt.Scanln(&i)

	fmt.Println("digite o número de valores k:")
	fmt.Scanln(&k)

	fmt.Println("digite o valor s:")
	fmt.Scanln(&s)

	values := make([]float64, int(k))
	for j := range values {
		values[j] = i + float64(j)*s
	}

	fmt.Println("tabuada de soma :")
	for _, B := range values {
		R := n + B
		fmt.Printf("%.2f + %.2f = %.2f\n", n, B, R)
	}

	fmt.Println("tabuada de subtração")
	for _, B := range values {
		R := B - n
		fmt.Printf("%.2f - %.2f = %.f\n", n, B, R)
	}

	fmt.Println("tabuada de multiplicação")
	for _, B := range values {
		R := n * B
		fmt.Printf("%.2f * %.2f = %.2f\n", n, B, R)
	}

	fmt.Println("Tabuada de divisão:")
	for _, B := range values {
		R := float64(n / B)
		fmt.Printf("%f / %f = %f\n", n, B, R)

	}
}
