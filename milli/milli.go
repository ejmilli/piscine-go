package main

import "github.com/01-edu/z01"

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}


func HextoRune(s byte) rune {

	if s < 10 {
		return rune(s + '0')
	}
	return rune(s-10+'a')
}

func PrintMemory(arr [10]byte) {
	  
	for i, char := range arr {
       firstDigit := char /16
       secondDigit := char %16

			 z01.PrintRune(HextoRune(firstDigit))
			 z01.PrintRune(HextoRune(secondDigit))



			 if (i+1)%4 == 0 || i == len(arr)-1 {
				z01.PrintRune('\n')
			 } else {
				z01.PrintRune(' ')
			 }
	}

	for _, char := range arr {
		if char >= 33 && char <= 126 {
			z01.PrintRune(rune(char))
		} else {
			z01.PrintRune('.')
		}
		
	}
	z01.PrintRune('\n')
}