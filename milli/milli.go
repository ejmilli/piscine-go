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
 
i := 0
var result string 

for i < len(s) {
	 count := 1
	 char := s[i]

	 for i+1 < len(s) && s[i+1] == char {
		count++
		i++
	 }
	 i++
	 result += strconv.Itoa(count) + string(char)
}
return result
}