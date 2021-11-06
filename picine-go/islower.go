package piscine

func IsLower(s string) bool {
	st := true
	for _, bcd := range s {
		if !(bcd >= 'a' && bcd <= 'z') {
			st = false
		}
	}
	return st
}
