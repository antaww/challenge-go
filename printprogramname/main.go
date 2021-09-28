package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	slashindex := -1
	for i, letter := range os.Args[0] {
		if letter == "/" {
			slashindex = i
		}
		for _, letter := range os.Args[0][slashindex+1:] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
