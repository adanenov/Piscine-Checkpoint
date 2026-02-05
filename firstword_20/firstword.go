package main

import (
	"fmt"
)

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}

func FirstWord(s string) string {
	if s == "" {
		return "\n"
	}
	a := ""
	for _, r := range s {
		if r == ' ' {
			break
		}
		a += string(r)
	}
	return a + "\n"
}
