package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if f(a[i], a[j]) == 1 {
				temp := a[i]
				modify(a, i, a[j])
				modify(a, j, temp)
			}
		}
	}
}

func modify(array []string, ind int, val string) {
	array[ind] = val
}
