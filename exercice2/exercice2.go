package main

import "fmt"

func CountLetters(texte string) int {
	i := 0
	for _, c := range texte {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			i++
		}
	}
	return i
}

func main() {

	fmt.Println(CountLetters("Hello,World"))
	fmt.Println(CountLetters("123 456"))
	fmt.Println(CountLetters("Golang.1.21"))
	fmt.Println(CountLetters("écolé"))
}
