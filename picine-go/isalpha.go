package piscine

func IsAlpha(s string) bool {
	st := true
	for _, bcd := range s {
		if !(bcd >= 'a' && bcd <= 'z' || bcd >= 'A' && bcd <= 'Z' || bcd >= '0' && bcd <= '9') {
			st = false
		}
	}
	return st
}
