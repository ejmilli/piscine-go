package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
if len(os.Args) != 2 {
	return
}

str := os.Args[1]
inWord := false
var result string

for _ , char := range str {
	if char > ' ' {
		result += string(char)
		inWord = true
	} else if inWord {
		result += "   "
		inWord = false
	}
}

if len(result) > 3 && result[len(result)-3:] == "   " {
	result = result[:len(result)-3]
}

for _, char := range result {
	z01.PrintRune(char)
}
z01.PrintRune('\n')

}