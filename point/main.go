package main

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		if n/10 != 0 {
			PrintNbr(n / -10)
		}
		mod := '0'
		for i := 0; i < -(n % 10); i++ {
			mod++
		}
		z01.PrintRune(mod)
	} else if n == 0 {
		z01.PrintRune('0')
	} else {
		if n/10 != 0 {
			PrintNbr(n / 10)
		}
		mod := '0'
		for i := 0; i < n%10; i++ {
			mod++
		}
		z01.PrintRune(mod)
	}
}

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)
	for _, letterA := range "x =" {
		z01.PrintRune(letterA)
	}
	PrintNbr(points.x)
	for _, letterB := range "y =" {
		z01.PrintRune(letterB)
	}
	PrintNbr(points.y)
	z01.PrintRune('\n')
}
