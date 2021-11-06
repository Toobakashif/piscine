package piscine

func UltimateDivMod(a *int, b *int) {
	var j int = 0
	var l int = 0
	j = *a
	l = *b
	*a = j / l
	*b = j % l
}
