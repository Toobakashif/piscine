package main

import (
	"fmt"
	"os"
)

var error bool

func main() {
	fileName := ""
	cnt := 0
	args := os.Args[1:]
	if args[0][0] == '-' {
		for _, c := range os.Args[2] {
			if c == '-' {
				os.Exit(1)
			}
			cnt = Atoi(os.Args[2])
		}
		args = args[2:]
	}
	for j, v := range args {
		fileName = v
		content, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("open " + fileName + ": no such file or directory")
			error = true
		} else {
			if j != 0 {
				fmt.Println()
			}
			fmt.Print("==> " + fileName + " <==\n")
			for i, v := range string(content) {
				if i >= len(string(content))-cnt {
					fmt.Print(string(v))
				}
			}
		}

	}
	if error {
		os.Exit(1)
	}
}

func Atoi(s string) int {
	er := -1
	m := 1
	r := []rune(s)
	result := 0
	for i := len(r) - 1; i > -1; i-- {
		if i == len(r)-1 {
			if int(r[i]) > 58 || int(r[i]) < 48 {
				er = 1
				break
			}
		}
		if int(r[i]) > 58 || int(r[i]) < 48 {
			if r[i] == '-' {
				result = -result
				er++
				continue
			}
			if r[i] == '+' {
				er++
				continue
			}
			er = 1
			break
		}
		result += (int(r[i]) - 48) * m
		m = m * 10
	}
	if er > 0 {
		return 0
	} else {
		return result
	}
}
