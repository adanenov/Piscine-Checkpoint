package main

import (
	"fmt"
)

func PrintIfNot(str string) string {
	if len(str) >= 3 {
		return "Invalid Input\n"
	}
	if str == "" {
		return "G\n"
	}
	for _, r := range str {
		if r >= '0' && r <= '9' {
			return "G\n"
		}
	}
	return "Invalid Input\n"
}

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}
