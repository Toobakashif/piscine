package piscine

func SortWordArr(a []string) {
	count := 0
	for c := range a {
		count = c + 1
	}

	temp := 0
	for i := 0; i < count-1; i++ {
		temp = i
		for j := i + 1; j < count; j++ {
			if a[j] < a[temp] {
				temp = j
			}
		}
		a[i], a[temp] = a[temp], a[i]
	}
}
