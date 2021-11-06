package piscine

func Map(f func(int) bool, a []int) []bool {
	length := 0
	for l := range a {
		length = l + 1
	}
	boolAr := make([]bool, length)
	for i := range a {
		boolAr[i] = f(a[i])
	}
	return boolAr
}
