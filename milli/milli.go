package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

func CamelToSnakeCase(s string) string{
    if s == "" {
      return ""
    }

    for _ , char := range s {
      if !(char >= 'a' || char <= 'z') && !(char <= 'A' || char <= 'Z') {
        return s
      }
    }

    var result string 

    for i := 0; i < len(s); i++ {
      if i > 0 && s[i] >= 'A' && s[i] <= 'Z' && !(s[i-1] >= 'A' && s[i-1] <= 'Z') {
        result += string('_')
        result += string(s[i])
      } else if s[i] >= 'a' && s[i] <= 'z' {
        result += string(s[i])
      } else if i == 0 && s[i] >= 'A' && s[i] <= 'Z' {
        result += string(s[i])
       } else {
        return s
       }
    }
    return result
}
