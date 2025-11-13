package main

import (
    "fmt"

)

func FirstWord(s string) string {
    if len(s) == 0 {
		return ""

	}
	
	for _,char:= range s {
		result:= string(char)

	
		if char == ' ' {
			
			
			break
		}
		return result
	}
	
	return result
}

func main() {
    fmt.Print(FirstWord("hello there"))
    // fmt.Print(FirstWord(""))
    // fmt.Print(FirstWord("hello   .........  bye"))
}