package main

import (
	"fmt"
	"os"
	"strconv"
)


func isFactor(n int)  {


	if n <= 1 {
		return
	}

	factor := 2

	for n > 1 {
		if n%factor == 0 {
			fmt.Print(factor)
			n/=factor
			if n > 1 {
				fmt.Print("*")
			} 
		} else {
			factor++
		}
	}
}

func main() {

	if len(os.Args) == 2 {
		 input , err := strconv.Atoi(os.Args[1]) 
		 if err == nil {
			isFactor(input)
		 }
		 fmt.Println()
	}
}