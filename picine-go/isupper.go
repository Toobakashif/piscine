package piscine

func IsUpper(s string) bool {
	pls := true
	for _, jh := range s {
		if !(jh >= 'A' && jh <= 'Z') {
			pls = false
		}
	}
	return pls
}
