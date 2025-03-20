package main

import (
	"fmt"
)

func main() {
	// Definir los meses con su cantidad de días
	meses := [3]string{"Junio", "Julio", "Agosto"}
	diasMes := [3]int{30, 31, 31} // Días de cada mes

	// Mostrar los meses en formato de calendario
	fmt.Println("Calendario de Junio, Julio y Agosto del año 2025")
	for i, mes := range meses {
		fmt.Println("\n", "        - ", mes, " -")
		fmt.Println("Lu  Ma  Mi  Ju  Vi  Sa  Do")

		// Simulación del inicio en lunes
		primerDia := 1 // Siempre comenzará en lunes

		// Imprimir espacios iniciales
		for j := 1; j < primerDia; j++ {
			fmt.Print("    ")
		}

		// Imprimir los días del mes
		for dia := 1; dia <= diasMes[i]; dia++ {
			fmt.Printf("%2d  ", dia)
			if (primerDia+dia-1)%7 == 0 { // Salto de línea cada domingo
				fmt.Println()
			}
		}
		fmt.Println("\n")
	}
}
