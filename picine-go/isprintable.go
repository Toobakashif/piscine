package piscine

func IsPrintable(s string) bool {
	sw := true
	sm := []rune(s)
	for _, key := range sm {
		if key >= 0 && key <= 31 {
			sw = false
		}
	}
	return sw
}
