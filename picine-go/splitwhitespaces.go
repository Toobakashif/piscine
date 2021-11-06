package piscine

func SplitWhiteSpaces(s string) []string {
	k := 0
	no_white_spaces := false
	for index := range s {
		if no_white_spaces && index != 0 && (s[index-1] == '\n' || s[index-1] == '\t' || s[index-1] == ' ') && s[index] != '\n' && s[index] != '\t' && s[index] != ' ' {
			k++
		}
		if s[index] != '\n' && s[index] != '\t' && s[index] != ' ' {
			no_white_spaces = true
		}
	}
	k++

	x := 0
	ans := make([]string, k)
	ok := true
	for _, c := range s {
		if c == '\n' || c == '\t' || c == ' ' {
			if !ok {
				x++
			}
			ok = true
			continue
		}
		ans[x] = ans[x] + string(c)
		ok = false
	}
	return ans
}
