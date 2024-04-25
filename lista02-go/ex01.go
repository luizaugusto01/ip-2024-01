package main

import (
	"fmt"
)

func main() {

	//matrícula
	var matricula float64
	fmt.Print(" digite a sua matrícula ")
	fmt.Scanln(&matricula)

	//notasprovas
	var notas [8]float64
	mnota := 0.0
	for i := 0; i < 8; i++ {
		fmt.Println("digite um a nota", i+1)
		fmt.Scanln(&notas[i])
		mnota = mnota + notas[i]
	}

	//notalistas
	var listas [5]float64
	mlista := 0.0
	for j := 0; j < 5; j++ {
		fmt.Println("digite as notas da lista", j+1)
		fmt.Scanln(&notas[j])
		mlista = mlista + listas[j]
	}

	//nota trabalho final
	var notatrabalhofinal float64
	fmt.Println("digite a nota do trabalho final")
	fmt.Scan(&notatrabalhofinal)

	//frequencia
	var freq int
	fmt.Println("digite sua frequencia")
	fmt.Scan(&freq)

	mediatotal := (0.7 * mnota / 8) + (0.15 * mlista / 5) + (0.15 * notatrabalhofinal)

	if mediatotal >= 6 && freq >= 96 {
		println("aprovado")
	} else if mediatotal >= 6 && freq < 96 {
		println("reprovado por falta")
	} else if mediatotal < 6 && freq >= 96 {
		println("reprovado por falta")
	} else {

		fmt.Println("reprovado por nota e falta")
	}

}
