package main

import (
	"fmt"
)

func CamelToSnakeCase(s string) string {
	for i, r := range s {
		if r >= '1' && r <= '9' {
			return s
		}
		if i > 0 && r >= 'A' && r <= 'Z' {
			prev := rune(s[i-1])
			if prev >= 'A' && prev <= 'Z' {
				return s
			}
		}
	}
	last := s[len(s)-1]
	if last >= 'A' && last <= 'Z' {
		return s
	}
	result := ""
	for i, r := range s {
		if r >= 'A' && r <= 'Z' && i != 0 {
			result += "_"
		}
		result += string(r)
	}
	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
