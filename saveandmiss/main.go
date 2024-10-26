package main

import (
	"fmt"
)

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg // If num is 0 or negative, return the full string
	}

	result := ""

	for i := 0; i < len(arg); i += 2 * num {
		end := i + num
		if end > len(arg) {
			end = len(arg) // Ensure we don't go out of bounds
		}

		// Concatenate the selected characters to the result string
		result += arg[i:end]
	}
	return result
}