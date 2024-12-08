package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	result := ""
	inWord := false

	for _, char := range input {
		//if char != ' ' && char != '\t' {
		if char > ' ' {
			result += string(char)
			inWord = true
		} else if inWord {
			// Only add exactly three spaces when transitioning out of a word
			result += "   "
			inWord = false
		}
	}

	// Remove any extra spaces at the end of result
	if len(result) > 3 && result[len(result)-3:] == "   " {
		result = result[:len(result)-3]
	}

	// Print the result
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}