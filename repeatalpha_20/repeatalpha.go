package main

import (
	"fmt"
)

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}

func RepeatAlpha(s string) string {
	result := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			count := int(r-'a') + 1
			for i := 0; i <= count; i++ {
				result += string(r)
			}
		} else if r >= 'A' && r <= 'Z' {
			count := int(r-'A') + 1
			for i := 0; i <= count; i++ {
				result += string(r)
			}
		} else {
			result += string(r)
		}
	}
	return result
}
