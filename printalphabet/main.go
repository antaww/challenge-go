package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var aRune string = "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 26; i++ {
		z01.PrintRune(rune(aRune[i]))
	}
	z01.PrintRune('\n')
}
