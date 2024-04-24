package main

import "fmt"

var salarioMinimo, kWh float64

func main() {

	// Leitura do valor do salário mínimo e quantidade de kWh
	fmt.Println("Digite o valor do salário mínimo:")
	fmt.Scanln(&salarioMinimo)
	fmt.Println("Digite a quantidade de kWh gasta pela residência:")
	fmt.Scanln(&kWh)

	// Calculo do custo por kWh
	custoPorKWh := (70.0 / 100.0) * salarioMinimo / 100.0

	// Calculo do custo do consumo da residência
	custoConsumo := kWh * custoPorKWh

	// Calculo do custo com desconto de 10%
	custoComDesconto := custoConsumo * 0.9

	// Impressão dos resultados
	fmt.Printf("Custo por kWh: R$ %.2f\n", custoPorKWh)
	fmt.Printf("Custo do consumo: R$ %.2f\n", custoConsumo)
	fmt.Printf("Custo com desconto: R$ %.2f\n", custoComDesconto)
}
