package main

import "fmt"

func CheckNumber(arg string) bool {
	for _, values := range arg {
		if values >= '0' && values <= '9' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}
