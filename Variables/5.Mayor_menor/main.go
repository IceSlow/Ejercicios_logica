// Escriba un programa que tome una lista de números como entrada y encuentre el número más grande y más pequeño en la lista
package main

import "fmt"

func main() {
	var Numeros [5]uint8
	fmt.Scanln(&Numeros[0])
	fmt.Scanln(&Numeros[1])
	fmt.Scanln(&Numeros[2])
	fmt.Scanln(&Numeros[3])
	fmt.Scanln(&Numeros[4])

	var (
		Numero_menor uint8
		Numero_mayor uint8
	)

	fmt.Print(Numero_menor, Numero_mayor)
}
