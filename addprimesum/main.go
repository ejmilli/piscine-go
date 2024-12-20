package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isSum(n int)  int {
	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

func isDigits(n int) {
	if n == 0 {
		return
	}

	isDigits(n / 10)
	z01.PrintRune(rune(n % 10 +'0'))
}

func printNumber(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	isDigits(n)
}

func main() {

	if len(os.Args) != 2 {

		z01.PrintRune('0')
		return
	}

	args := os.Args[1]
	n := 0

	for _, char := range args {
		if char < '0' || char > '9' {
			z01.PrintRune('0')
			return
		}

		n = n*10 + int(char-'0')
	}

	result := isSum(n)
	printNumber(result)
	z01.PrintRune('\n')
}
