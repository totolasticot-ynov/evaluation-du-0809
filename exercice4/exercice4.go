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

	fmt.Print("first : ")
	fmt.Scan(&x)
	fmt.Print("second : ")
	fmt.Scan(&y)

	fmt.Printf(" le plus grand diviseur commun de %d et %d est %d\n", x, y, pgcd(x, y))
}
