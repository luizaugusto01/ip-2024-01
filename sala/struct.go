package main

import (
	"fmt"
)

// Definindo a struct Pessoa
type Pessoa struct {
	Nome   string
	Altura float64
	Peso   float64
}

// Função para calcular o peso ideal
func calcularPesoIdeal(altura float64) float64 {
	return 72.7*altura - 58.0
}

func main() {
	var pessoas []Pessoa

	// Loop para ler os dados até que "FIM" seja digitado como nome
	for {
		var nome string
		fmt.Print("Digite o nome da pessoa (ou 'FIM' para sair): ")
		fmt.Scanln(&nome)

		if nome == "FIM" {
			break
		}

		var altura float64
		fmt.Print("Digite a altura da pessoa (em metros): ")
		fmt.Scanln(&altura)

		// Calculando o peso ideal
		pesoIdeal := calcularPesoIdeal(altura)

		// Criando uma nova pessoa com os dados fornecidos
		pessoa := Pessoa{
			Nome:   nome,
			Altura: altura,
			Peso:   pesoIdeal,
		}

		// Adicionando a nova pessoa ao slice de pessoas
		pessoas = append(pessoas, pessoa)
	}

	// Imprimindo o slice de pessoas
	fmt.Println("Lista de Pessoas:")
	for _, pessoa := range pessoas {
		fmt.Printf("Nome: %s, Altura: %.2fm, Peso Ideal: %.2fkg\n", pessoa.Nome, pessoa.Altura, pessoa.Peso)
	}
}