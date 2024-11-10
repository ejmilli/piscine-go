package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) != 4 {
		return
	}

	str := os.Args[1]
	old := os.Args[2]
	new := os.Args[3]


	if len(old) != 1 && len(new) !=  1 {
		return
	}
 var result string

	for _, char := range str {
		if string(char) == old {
			result += new 
		} else {
			result += string(char)
		}
	}

	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')

}