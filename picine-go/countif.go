package piscine

func CountIf(f func(string) bool, tab []string) int {
	c := 0
	for _, ch := range tab {
		if f(ch) {
			c++
		}
	}
	return c
}
