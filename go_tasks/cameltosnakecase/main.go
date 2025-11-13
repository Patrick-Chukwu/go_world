package main

import (
	"fmt"
)

func CamelToSnakeCase(s string) string{

	result := ""
	if len(s) == 0 {
		return ""
	}
	for i, value := range s {
		if (i == len(s) - 1 && s[i] >= 'A' && s[i] <= 'Z') {
		return s
	} else if value >= '0' && value <= '9' {
		return s
	}

	if i == 0  && value >= 'A' && value <= 'Z' {
		
	}

	
	if i != 0 && value >= 'A' && value <= 'Z' {
		result += "_"
		

	} 
	result += string(value)

	}
	return result



}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase(""))
	fmt.Println(CamelToSnakeCase("hey2"))

}