package main

import "github.com/01-edu/z01"

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func HextoRune(s byte) rune {
	if s < 10 {
		return rune(s + '0')
	}
	return rune(s - 10 + 'a')
}

func PrintMemory(arr [10]byte) {
	for i, char := range arr {
		firstDigit := char / 16
		secondDigit := char % 16

		z01.PrintRune(HextoRune(firstDigit))
		z01.PrintRune(HextoRune(secondDigit))
		z01.PrintRune(' ')

		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		}
	}
	z01.PrintRune('\n')

	for _, char := range arr {
		if char >= 32 && char <= 126 {
			z01.PrintRune(rune(char))
		} else {
			z01.PrintRune('.')
		}

	}
	z01.PrintRune('\n')

}