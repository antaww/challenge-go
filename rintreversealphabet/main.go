package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var aRune string = "abcdefghijklmnopqrstuvwxyz" //string a lire
	for i := 26; i < 1; i-- {                       //boucle 26 repp
		z01.PrintRune(rune(aRune[i])) //rune valeur i pour decrementer
	}
	z01.PrintRune('\n')
}
