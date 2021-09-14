package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var aRune string = "0123456789"
	for i := 0; i < 10; i++ {
		z01.PrintRune(rune(aRune[i]))
	}
	z01.PrintRune('\n')
}
