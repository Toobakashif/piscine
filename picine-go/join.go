package piscine

func Join(strs []string, sep string) string {
	fr := ""
	first := true
	for _, fra := range strs {
		if first {
			fr = fra
			first = false

		} else {
			fr += sep + fra
		}
	}
	return fr
}
