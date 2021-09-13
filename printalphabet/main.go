package main

import (
	"github.com/01-edu/z01"
)

func main() {

	var aRune string = "abcdefghijklmnopqrstuvwxyz'\n'" //string a lire
	for i := 0; i < 26; i++ {                           //boucle 26 rep
		z01.PrintRune(rune(aRune[i])) //rune valeur i pour incrementer
	}
}
