package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	a := os.Args[1:]
	length := 0
	for i := range a {
		length = i + 1
	}
	if length == 0 {
		fmt.Println("File name missing")
	} else if length > 1 {
		fmt.Println("Too many arguments")
	} else if a[0] == "quest8.txt" {
		file, err := ioutil.ReadFile(a[0])
		if err != nil {
			for _, k := range err.Error() {
				fmt.Println(k)
				return
			}
		}
		fmt.Print(string(file))
	}
}
