package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	myargs := os.Args[1:]
	if len(myargs) < 1 {
		z01.PrintRune('\n')
		return
	}
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	var vowelsfound []rune
	for i := 0; i < len(myargs); i++ {
		for j := 0; j < len(myargs[i]); j++ {
			for x := 0; x < len(vowels); x++ {
				if rune(myargs[i][j]) == vowels[x] || rune(myargs[i][j]+32) ==
					vowels[x] {
					vowelsfound = append(vowelsfound, rune(myargs[i][j]))
				}
			}
		}
	}
	for _, sona := range myargs {
		for _, taht := range sona {
			if taht == 'a' || taht == 'e' || taht == 'i' || taht == 'o' || taht == 'u' ||
				taht == 'A' || taht == 'E' || taht == 'I' || taht == 'O' || taht == 'U' {
				taht = vowelsfound[len(vowelsfound)-1]
				vowelsfound = vowelsfound[:len(vowelsfound)-1]
			}
			z01.PrintRune(taht)
		}
		z01.PrintRune(rune(32))
	}
	z01.PrintRune('\n')
}
