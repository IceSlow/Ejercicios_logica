package main

import "fmt"

var (
	numero_1 float64
	numero_2 float64
	decision uint8
)

func elegir_numeros() {
	fmt.Println("Escribe los dos numeros que deseas operar")
	fmt.Scanln(&numero_1)
	fmt.Scanln(&numero_2)
}

func operacion() {
	if decision == 1 {
		fmt.Print("La suma es: ", numero_1+numero_2)
	} else if decision == 2 {
		fmt.Print("La resta es: ", numero_1-numero_2)
	} else if decision == 3 {
		fmt.Println("La multiplicación es: ", numero_1*numero_2)
	} else if decision == 4 {
		fmt.Println("Su división es: ", numero_1/numero_2)
	} else {
		fmt.Println("ERROR ", decision, " no es valido")
	}
}

func main() {
	fmt.Println("Bienvenido a la calculadora!")
	fmt.Printf(" Para Sumar: 1\n Para Restar 2\n Para multiplicar 3\n Para dividir 4\n")
	fmt.Scanln(&decision)
	elegir_numeros()
	operacion()
}
