package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for i := '0'; i <= '8'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for k := j + 1; k <= '8'; k++ {
				for l := k + 1; l <= '8'; l++ {
					if j <= 9 {
						i := i + 1
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(32)
						z01.PrintRune(k)
						z01.PrintRune(l)
						if k == '9' && l == '9' {
							j := j + 1
							k := k - 9
							l := l - 9
							if i == '9' && j == '8' && k == '9' && l == '9' {
								z01.PrintRune(10)
							} else {
								z01.PrintRune(44)
								z01.PrintRune(32)
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
