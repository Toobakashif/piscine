package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	len := 0
	for i := range a {
		len = i + 1
	}

	asc := true
	desc := true

	for i := 1; i < len; i++ {
		if !(f(a[i-1], a[i]) <= 0) {
			asc = false
		}
	}

	for i := 1; i < len; i++ {
		if !(f(a[i-1], a[i]) >= 0) {
			desc = false
		}
	}

	return asc || desc
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
