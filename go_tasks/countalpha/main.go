package main

import (
	"fmt"

)
func CountAlpha(s string) int {

	number := 0

	for _, ch := range s {
		
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		 number++
		}
		
	}
	return number

}

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}
