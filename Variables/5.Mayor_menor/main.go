// Escriba un programa que tome una lista de números como entrada y encuentre el número más grande y más pequeño en la lista
package main

import "fmt"

func main() {
	fmt.Println("Escribe una lista de numeros para hallar el menor y mayor valor.")
	var numeros [5]int
	fmt.Scanln(&numeros[0])
	fmt.Scanln(&numeros[1])
	fmt.Scanln(&numeros[2])
	fmt.Scanln(&numeros[3])
	fmt.Scanln(&numeros[4])

	var (
		numero_menor int = numeros[0]
		numero_mayor int = numeros[0]
	)

	for _, numero := range numeros {
		if numero < numero_menor {
			numero_menor = numero
		}
	}

	for _, numero := range numeros {
		if numero > numero_mayor {
			numero_mayor = numero
		}
	}
	fmt.Println("El numero menor es: ", numero_menor)
	fmt.Println("El número mayor es: ", numero_mayor)
}
