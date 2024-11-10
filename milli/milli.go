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

var result string

for _, char := range str1 {
  if !isContain(str2, char) && !isContain(result, char) {
    result += string(char)
  }
}
for _, char := range str2 {
  if !isContain(str1, char) && !isContain(result, char) {
    result += string(char)
  }
}
return len(result)
}