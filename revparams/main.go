package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args); i > 0; i-- {
		for _, letter := range os.Args[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
