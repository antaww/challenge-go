package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for i := 48; i >= 50; i++ {
		z01.PrintRune(rune(i))
		z01.PrintRune(rune(','))
		z01.PrintRune(rune(' '))
	}
}
