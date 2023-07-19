// Cree un programa que tome un número como entrada y muestre la tabla de multiplicar de ese número hasta el 10.
package main

import "fmt"

func main() {
	var (
		number_user uint8
	)
	fmt.Println("Escriba un numero entero:")
	fmt.Scanln(&number_user)
	fmt.Println("La tabla del multiplicar del 1 al 12 de: ", number_user)
	fmt.Println(number_user, "x 1 =", number_user*1)
	fmt.Println(number_user, "x 2 =", number_user*2)
	fmt.Println(number_user, "x 3 =", number_user*3)
	fmt.Println(number_user, "x 4 =", number_user*4)
	fmt.Println(number_user, "x 5 =", number_user*5)
	fmt.Println(number_user, "x 6 =", number_user*6)
	fmt.Println(number_user, "x 7 =", number_user*7)
	fmt.Println(number_user, "x 8 =", number_user*8)
	fmt.Println(number_user, "x 9 =", number_user*9)
	fmt.Println(number_user, "x 10 =", number_user*10)
	fmt.Println(number_user, "x 11 =", number_user*11)
	fmt.Println(number_user, "x 12 =", number_user*12)

}
