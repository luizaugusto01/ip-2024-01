package main

import (
	"fmt"
	"strings"
)

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var cpf string
		fmt.Scan(&cpf)

		if validarCPF(cpf) {
			fmt.Println("CPF valido")
		} else {
			fmt.Println("CPF invalido")
		}
	}
}

func validarCPF(cpf string) bool {
	// Remover os espaços em branco da string do CPF
	cpf = strings.ReplaceAll(cpf, " ", "")

	// Verificar se o CPF tem 11 dígitos
	if len(cpf) != 11 {
		return false
	}

	// Verificar se todos os caracteres do CPF são dígitos
	for _, char := range cpf {
		if char < '0' || char > '9' {
			return false
		}
	}

	// Verificar o primeiro dígito verificador
	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(cpf[i]-'0') * (10 - i)
	}
	remainder := sum % 11
	digit1 := 11 - remainder
	if digit1 >= 10 {
		digit1 = 0
	}

	// Verificar o segundo dígito verificador
	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(cpf[i]-'0') * (11 - i)
	}
	remainder = sum % 11
	digit2 := 11 - remainder
	if digit2 >= 10 {
		digit2 = 0
	}

	// Verificar se os dígitos verificadores correspondem aos do CPF
	return digit1 == int(cpf[9]-'0') && digit2 == int(cpf[10]-'0')
}
