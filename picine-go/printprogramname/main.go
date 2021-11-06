package main

import (
	"os"

	"github.com/01-edu/z01"
)

func WriteName(s string) {
	x := []rune(s)
	for i := FindIndex(x, rune('/')) + 1; i < len(x); i++ {
		z01.PrintRune(x[i])
	}
	z01.PrintRune(rune('\n'))
}

func main() {
	ar := os.Args
	WriteName(ar[0])
}

func FindIndex(runesslice []rune, find rune) int {
	for i := len(runesslice) - 1; i >= 0; i-- {
		if runesslice[i] == find {
			return i
		}
	}
	return 0
}
