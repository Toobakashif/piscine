package piscine

func BasicAtoi2(s string) int {
	r := []rune(s)
	Ans := 0
	for i := 0; i < len(s); i++ {
		if int(r[i]) < 48 || int(r[i]) > 57 {
			return 0
		} else {
			Ans *= 10
			Ans += int(r[i]) - '0'
		}
	}
	return Ans
}
