package main

import (
	"fmt"
)

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}

func RepeatAlpha(s string) string {

  var result string 

  for _, char := range s {
    if char >= 'a' && char <= 'z' {
      for i := 0; i < int((char-'a')+1); i++ {
        result += string(char)
      }
    } else if char >= 'A' && char <= 'Z' {
      for i := 0; i < int((char -'A')+1); i++ {
        result += string(char)
      }
    } else {
      result += string(char)
    }
  }
  return result
}