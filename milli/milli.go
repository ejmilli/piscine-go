package main

import (
	"fmt"
)

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}

func LastWord(s string) string {
	var lastWord string 
	var currentWord string 

	for _, char := range s {
		if char == ' ' {
			if currentWord != "" {
				lastWord = currentWord
			}
			currentWord  = ""
		} else {
			currentWord += string(char)
		}
	}

	if currentWord != "" {
		lastWord = currentWord
	}

	return lastWord + "\n"
}