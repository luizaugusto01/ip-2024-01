package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		var matricula int
		var horastrabalhadas, valorporhora float64
		fmt.Println("me passe os dados :")
		_, err := fmt.Scanf("%d %f %f\n", &matricula, &horastrabalhadas, &valorporhora)
		if err != nil {
			fmt.Println("Erro ao ler a entrada:", err)
			return
		}
		if matricula == 0 {
			break
		}
		salario := horastrabalhadas * valorporhora

		fmt.Printf("%d %.2f\n", matricula, salario)

		bufio.NewReader(os.Stdin).ReadByte()
	}

}
