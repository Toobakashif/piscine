package main

import (
	"os"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) < 3 {
		return
	}
	if !CheckIfNumber(arguments[0]) || !CheckIfNumber(arguments[2]) {
		return
	}
	x := ToInt(arguments[0])
	y := ToInt(arguments[2])

	// OVERFLOW CHECKS
	if len(arguments[0]) >= 19 || len(arguments[2]) >= 19 {
		return
	}
	if x >= 9223372036854775807 || x <= -9223372036854775808 {
		return
	}
	if y >= 9223372036854775807 || y <= -9223372036854775808 {
		return
	}
	//
	op := (arguments[1])
	result := 0
	if (x == 0 || y == 0) && (op == "/" || op == "%") {
		if op == "/" {
			os.Stdout.WriteString("No division by 0")
			os.Stdout.WriteString("\n")
			return
		}
		if op == "%" {
			os.Stdout.WriteString("No modulo by 0")
			os.Stdout.WriteString("\n")
			return
		}
	} else if op == "%" || op == "/" || op == "*" || op == "+" || op == "-" {
		result = calculate(x, y, op)
	} else {
		return
	}
	strResult := toStr(result)
	os.Stdout.WriteString(strResult)
	os.Stdout.WriteString("\n")
}

func calculate(a, b int, op string) (result int) {
	if op == "+" {
		return a + b
	} else if op == "-" {
		return a - b
	} else if op == "*" {
		return a * b
	} else if op == "/" {
		return a / b
	} else if op == "%" {
		return a % b
	} else {
		return
	}
}

func ToInt(s string) (result int) {
	bites := []byte(s)
	isNeg := false
	for _, c := range bites {
		if c == 45 {
			isNeg = true
		}
		if 48 <= c && c <= 57 {
			result *= 10
			result += int(c - 48)
		}
	}
	if isNeg {
		result *= -1
	}
	return
}

func toStr(nr int) string {
	str := ""
	isNeg := false
	if nr == 0 {
		return "0"
	}
	if nr < 0 {
		nr *= -1
		isNeg = true
	}
	for nr != 0 {
		// fmt.Print(str)
		str = string(rune(nr%10)+48) + str
		nr /= 10
	}
	m := ""
	if isNeg {
		m = "-" + str
		return m
	}
	return str
}

func CheckIfNumber(s string) (isNr bool) {
	bites := []byte(s)
	for _, i := range bites {
		if (47 <= i && i <= 57) || i == 45 {
			// fmt.Print(i)
			isNr = true
		} else {
			isNr = false
		}
	}
	return
}
