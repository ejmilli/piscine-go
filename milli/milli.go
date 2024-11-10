package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func ZipString(s string) string {

var result string 
i := 0 

for i < len(s) {
  char := s[i]
  count := 1
  for i+1 < len(s) && s[i+1] == char {
    count++
    i++
  }
  result += strconv.Itoa(count) + string(char)
  i++
}
return result
}