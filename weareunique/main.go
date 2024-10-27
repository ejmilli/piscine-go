package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}

func isContain(s string, r rune) bool {
	for _, char := range s {
		if char == r {
			return true
		}
	}
	return false
}

func WeAreUnique(str1, str2 string) int {

	if len(str1) == 0 && len(str2) == 0 {
		return -1
	}

	if str1 == str2 {
		return 0
	}

	result := ""

	for _, char1 := range str1 {
		if !isContain(str2, char1) && !isContain(result, char1) {
			result += string(char1)
		}

		for _, char2 := range str2 {
			if !isContain(str1, char2) && !isContain(result, char2) {
				result += string(char2)
			}
		}

	}
	return len(result)
}
