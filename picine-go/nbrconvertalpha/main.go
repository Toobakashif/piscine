package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	i := 1
	a := 'a'

	if len(os.Args) < 2 {
		return
	}

	if os.Args[1] == "--upper" {
		i = 2
		a = 'A'
	}

	for ; i < len(os.Args); i++ {
		n := Atoi(os.Args[i])
		if 0 < n && n < 27 {
			z01.PrintRune(rune(n-1) + a)
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	sArr := []rune(s)
	minus := 1
	if sArr[0] == '-' {
		minus = -1
		sArr[0] = '0'
	} else if sArr[0] == '+' {
		sArr[0] = '0'
	}
	n := 0
	for i := 0; i < len(sArr); i++ {
		if sArr[i] < '0' || '9' < sArr[i] {
			return 0
		}
		n *= 10
		n += int(sArr[i] - '0')
	}
	return minus * n
}
