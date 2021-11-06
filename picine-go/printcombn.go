// package main
package piscine

import "github.com/01-edu/z01"

func IntToString(num int) string {
	switch num {
	case 1:
		return "1"
	case 2:
		return "2"
	case 3:
		return "3"
	case 4:
		return "4"
	case 5:
		return "5"
	case 6:
		return "6"
	case 7:
		return "7"
	case 8:
		return "8"
	case 9:
		return "9"
	default:
		return "0"
	}
}

func itoa(num int) string {
	var result string
	if num == 0 {
		return "0"
	}

	for num > 0 {
		result = IntToString(num%10) + result
		num /= 10
	}
	return result
}

func printString(str string) {
	for _, c := range str {
		z01.PrintRune(c)
	}
}

func printComb(n int, prev int, result string, count *int) {
	for i := 0; i < 10; i++ {
		if prev < i {
			if n == 1 {
				if *count > 0 {
					printString(", ")
				}
				printString(result + itoa(i))
				*count++
			} else {
				printComb(n-1, i, result+itoa(i), count)
			}
		}
	}
}

func PrintCombN(n int) {
	var count int = 0
	for i := 0; i < 10; i++ {
		if n > 1 {
			printComb(n-1, i, itoa(i), &count)
		} else {
			if count > 0 {
				printString(", ")
			}
			printString(itoa(i))
			count++
		}
	}
	printString("\n")
}

/*func main(){
	PrintCombN(1)
}
*/
