package main

import (
	"fmt"
)

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}

func ConcatAlternate(s1, s2 []int) []int {

	var result []int

	if len(s1) >= len(s2) {
		for i := 0; i < len(s1); i++ {
			result = append(result, s1[i])
			if i < len(s2) {
				result = append(result, s2[i])
			}
		}
	} else if len(s2) > len(s1) {
		for i := 0; i < len(s2); i++ {
			result = append(result, s2[i])
			if i < len(s1) {
				result = append(result, s1[i])
			}
		}
	} 
return result
}