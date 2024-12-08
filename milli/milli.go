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


	var firstWord string 
	inWord := false

	for _ , char := range s {
		if char != ' ' {
			firstWord += string(char)
			inWord = true
		} else if inWord {
			break
		}
	}
	return firstWord + "\n"

}