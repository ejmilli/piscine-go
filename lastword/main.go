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
	foundWord := false

	for _, char := range s {
		if char == ' ' {
			if !foundWord {
				continue
			}

			if foundWord {
				break

			}

		}
		lastWord = lastWord + string(char)
		foundWord = true

	}
	return lastWord + "\n"
}