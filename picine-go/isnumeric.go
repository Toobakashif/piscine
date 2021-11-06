package piscine

func IsNumeric(s string) bool {
	ss := true
	for _, g := range s {
		if !(g >= '0' && g <= '9') {
			ss = false
		}
	}
	return ss
}
