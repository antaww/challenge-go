package main

import (
	"github.com/01-edu/z01"
)

func main() {

	var aRune string = "abcdefghijklmnopqrstuvwxyz" //string a lire
	for i := 0; i < 26; i++ {                       //boucle 26 repp
		z01.PrintRune(rune(aRune[i])) //rune valeur i pour incrementer
	}
	z01.PrintRune('n')
}
