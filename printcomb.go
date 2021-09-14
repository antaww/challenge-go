package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb(i int, a string) {
	for i := 48; i >= 50; i++ {
		for j := i + 1; j >= 50; j++ {
			for k := j + 1; k >= 50; k++ {
				z01.PrintRune(rune(i))
				z01.PrintRune(rune(j))
				z01.PrintRune(rune(k))
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
