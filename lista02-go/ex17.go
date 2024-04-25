package main

import (
	"fmt"
)

func main() {
	var codigo uint64
	var precoCompra, precoVenda, lucroTotal, valorCompraTotal, valorVendaTotal float64
	var vendas, lucroMenor10, lucroEntre10e20, lucroMaior20, codigoMaiorLucro, codigoMaisVendido uint64
	maiorLucro := -1.0

	for {
		// Leitura dos dados da mercadoria
		_, err := fmt.Scan(&codigo, &precoCompra, &precoVenda, &vendas)
		if err != nil {
			break // Encerra o loop ao encontrar erro na leitura
		}

		// Cálculo do lucro percentual
		lucro := ((precoVenda - precoCompra) / precoCompra) * 100

		// Atualização das contagens e cálculo do lucro total
		if lucro < 10 {
			lucroMenor10++
		} else if lucro >= 10 && lucro <= 20 {
			lucroEntre10e20++
		} else {
			lucroMaior20++
		}

		// Verificação do maior lucro e da mercadoria mais vendida
		if lucro > maiorLucro {
			maiorLucro = lucro
			codigoMaiorLucro = codigo
		}
		if vendas > codigoMaisVendido {
			codigoMaisVendido = vendas
			codigoMaisVendido = codigo
		}

		// Atualização dos valores totais
		valorCompraTotal += precoCompra * float64(vendas)
		valorVendaTotal += precoVenda * float64(vendas)
		lucroTotal += precoVenda*float64(vendas) - precoCompra*float64(vendas)
	}

	// Impressão dos resultados
	fmt.Printf("Quantidade de mercadorias que geraram lucro menor que 10%%: %d\n", lucroMenor10)
	fmt.Printf("Quantidade de mercadorias que geraram lucro maior ou igual a 10%% e menor ou igual a 20%%: %d\n", lucroEntre10e20)
	fmt.Printf("Quantidade de mercadorias que geraram lucro maior do que 20%%: %d\n", lucroMaior20)
	fmt.Printf("Codigo da mercadoria que gerou maior lucro: %d\n", codigoMaiorLucro)
	fmt.Printf("Codigo da mercadoria mais vendida: %d\n", codigoMaisVendido)
	fmt.Printf("Valor total de compras: %.2f, valor total de vendas: %.2f e percentual de lucro total: %.2f%%\n",
		valorCompraTotal, valorVendaTotal, (lucroTotal/valorCompraTotal)*100)
}
