package piscine

func DivMod(a int, b int, div *int, mod *int) {
	*div = *div + a/b
	*mod = *mod + a%b
}
