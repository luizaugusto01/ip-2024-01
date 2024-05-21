package main

import (
	"fmt"
)

func main() {
	var (
		codigo                                          uint64
		precoCompra, precoVenda                         float64
		qtdVendida                                      int
		menor10, entre10e20, maior20                    uint64
		totalCompra, totalVenda, lucroTotal, maiorLucro float64
		codMaiorLucro, codMaisVendido                   uint64
		maiorQtdVendida                                 int
	)

	for {
		// Lendo os dados da mercadoria
		_, err := fmt.Scan(&codigo, &precoCompra, &precoVenda, &qtdVendida)
		if err != nil {
			break
		}

		// Calculando o lucro e atualizando as variáveis conforme necessário
		lucro := (precoVenda - precoCompra) * float64(qtdVendida)
		totalCompra += precoCompra * float64(qtdVendida)
		totalVenda += precoVenda * float64(qtdVendida)
		lucroTotal += lucro

		// Atualizando as contagens de lucro
		if lucro < precoCompra*float64(qtdVendida)*0.10 {
			menor10++
		} else if lucro <= precoCompra*float64(qtdVendida)*0.20 {
			entre10e20++
		} else {
			maior20++
		}

		// Verificar se esse é o maior lucro até agora
		if lucro > maiorLucro {
			maiorLucro = lucro
			codMaiorLucro = codigo
		}

		// Verificar se essa é a maior quantidade vendida até agora
		if qtdVendida > maiorQtdVendida {
			maiorQtdVendida = qtdVendida
			codMaisVendido = codigo
		}
	}

	// Calculando o percentual de lucro total
	percentualLucroTotal := (lucroTotal / totalCompra) * 100

	// Imprimindo os resultados
	fmt.Printf("Quantidade de mercadorias que geraram lucro menor que 10%%: %d\n", menor10)
	fmt.Printf("Quantidade de mercadorias que geraram lucro maior ou igual a 10%% e menor ou igual a 20%%: %d\n", entre10e20)
	fmt.Printf("Quantidade de mercadorias que geraram lucro maior do que 20%%: %d\n", maior20)
	fmt.Printf("Codigo da mercadoria que gerou maior lucro: %d\n", codMaiorLucro)
	fmt.Printf("Codigo da mercadoria mais vendida: %d\n", codMaisVendido)
	fmt.Printf("Valor total de compras: %.2f, valor total de vendas: %.2f e percentual de lucro total: %.2f%%\n", totalCompra, totalVenda, percentualLucroTotal)

	// Imprimindo o caractere de quebra de linha final
	fmt.Println()
}
