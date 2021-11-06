package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, v := range a {
		PRune(v)
	}
}

func PRune(str string) {
	strRune := []rune(str)

	for _, v := range strRune {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
