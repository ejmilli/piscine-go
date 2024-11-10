package main

import (
	"fmt"
)

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}

func ThirdTimeIsACharm(str string) string {

if len(str) == 0 {
  return "\n"
}

if len(str) < 3 {
  return "\n"
}

var result string

for i := 2; i < len(str); i+=3 {
  result += string(str[i])
}
return result + "\n"
}