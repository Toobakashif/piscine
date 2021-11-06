package piscine

func isprime(nb int) bool {
	if nb < 2 {
		return false
	} else if 2 == nb || 3 == nb {
		return true
	} else {
		limit := nb / 2
		for i := 2; i <= int(limit); i++ {
			if nb%i == 0 {
				return false
			}
		}
		return true
	}
}

func FindNextPrime(nb int) int {
	if isprime(nb) {
		return nb
	} else {
		n := nb + 1
		for isprime(n) == false {
			n++
		}
		return n
	}
}
