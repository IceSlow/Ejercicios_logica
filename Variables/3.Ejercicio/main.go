package main

import (
	"fmt"
)

var (
	four uint8 = 78
	two  uint8 = 2
)

func main() {
	var dezplazar_binario = four << two
	fmt.Printf("Binario de 78 y 2: %b y %b\n", four, two)
	fmt.Println(dezplazar_binario)
	fmt.Printf("Binario: %b", dezplazar_binario)
}
