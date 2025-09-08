package main

import (
	"fmt"
)

func pgcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func main() {
	var x, y int

	fmt.Print("le premier chiffres ou nombres est : ")
	fmt.Scan(&x)
	fmt.Print("le deuxième chiffres ou nombres est : ")
	fmt.Scan(&y)

	fmt.Printf(" le plus grand diviseur commun de %d et %d est %d\n", x, y, pgcd(x, y))
}
