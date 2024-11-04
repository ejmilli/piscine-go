
package main

import (
	"os"

	"github.com/01-edu/z01"
)
 

func main() {

 if len(os.Args) != 4 {
  return 
}

str1 := os.Args[1]
str2 := os.Args[2]
str3 := os.Args[3]

if len(str2) != 1 && len(str3) != 1 {
 return
}

var result string 

for _, char := range str1 {
   if string(char) == str2 {
   result += str3
} else {
  result += string(char)
}
}

for _, char := range result {
  z01.PrintRune(char)
}

}