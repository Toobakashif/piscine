package main

import (
	"github.com/01-edu/z01"
)

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

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	IntoRune(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	IntoRune(points.y)
	z01.PrintRune('\n')
}

func IntoRune(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	Check(n)
}

func Check(s int) {
	k := s
	l := k
	factor := 1

	for l > 0 {
		l = l / 10
		factor = factor * 10
	}

	for factor > 1 {
		factor = factor / 10
		z01.PrintRune(Atoi(k / factor))
		k = k % factor
	}
}

func Atoi(nb int) rune {
	var a rune = '0'
	switch nb {

	case 0:
		a = '0'

	case 1:
		a = '1'

	case 2:
		a = '2'
	case 3:
		a = '3'
	case 4:
		a = '4'
	case 5:
		a = '5'
	case 6:
		a = '6'
	case 7:
		a = '7'
	case 8:
		a = '8'
	case 9:
		a = '9'
	}
	return a
}
