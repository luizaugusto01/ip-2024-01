package main

import (
	"fmt"
)

func main() {

	var numeros []float64
	i := 0
	for {
		var num float64

		i++
		fmt.Println("digite o número ", i)
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		numeros = append(numeros, num)
	}
	var somapar, somaimpar float64
	var countpar, countimpar int

	for _, num := range numeros {
		if int(num)%2 == 0 {
			somapar += num
			countpar++

		} else {
			somaimpar += num
			countimpar++
		}

	}
	mediapar := somapar / float64(countpar)
	mediaimpar := somaimpar / float64(countimpar)

	fmt.Printf("media par = %.2f\n", mediapar)
	fmt.Printf("media impar = %.2f\n", mediaimpar)

}
