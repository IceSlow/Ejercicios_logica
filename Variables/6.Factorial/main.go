package main

import "fmt"

var (
	number    uint64
	factorial uint64 = 1
)

func main() {
	fmt.Println("Calcula el factorial del numero que desees, escribalo:")
	fmt.Scanln(&number)
	for i := uint64(1); i <= number; i++ {
		factorial *= i
	}
	fmt.Print("El factorial de ", number, "! es: ", factorial)
}
