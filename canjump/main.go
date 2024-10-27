package main

import (
	"fmt"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1)) // Should print: true

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2)) // Should print: false

	input3 := []uint{0}
	fmt.Println(CanJump(input3))  // Should print: true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CanJump(steps []uint) bool { // Changed 'unit' to '[]uint'
	if len(steps) == 0 {
		return false // Fixed typo from 'retrun' to 'return'
	}

	maxjump := 0
	for i := 0; i < len(steps); i++ {
		if i > maxjump {
			return false // Fixed typo, added 'return' keyword
		}

		maxjump = max(maxjump, i+int(steps[i]))
		if maxjump >= len(steps)-1 { // Changed to '>=' to allow reaching the last index
			return true
		}
	}
	return false
}
