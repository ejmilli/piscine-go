package main

import (
	"fmt"
)

func main() {
    fmt.Print(FirstWord("hello there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("hello   .........  bye"))
}

func FirstWord(s string) string {
	firstWord := ""
	inWord := false

	for _, char := range s {
		if char != ' ' { // If the character is not a space
			firstWord += string(char) // Add character to firstWord
			inWord = true
		} else if inWord { // If we encounter a space after starting a word
			break // Stop processing further
		}
	}

	return firstWord + "\n"
}