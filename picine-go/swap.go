package piscine

func Swap(a *int, b *int) {
	tempx := *a
	tempy := *b
	*a = tempy
	*b = tempx
}
