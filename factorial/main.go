package main

import "fmt"

func Factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {

	fmt.Println(Factorial(5))
	fmt.Println(Factorial(4))
	fmt.Println(Factorial(8))

}