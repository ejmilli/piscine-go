package main

import (
	"os"

	"github.com/01-edu/z01"
)

func FormString(str1, str2 string) bool {
	i := 0
	j := 0
	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			i++
		}
		j++
	}
	return i == len(str1)
}

func main() {
	args := os.Args

	if len(os.Args) != 3 {
		return
	}

	str1 := args[1]
	str2 := args[2]

	if FormString(str1, str2) {
		for _, char := range str1 {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}