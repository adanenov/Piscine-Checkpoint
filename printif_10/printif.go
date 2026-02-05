package main

import (
	"fmt"
)

func PrintIf(str string) string {
	if str == "" {
		return "G\n"
	}
	for _, r := range str {
		if len(str) >= 3 && r >= 'a' && r <= 'z' {
			return "G\n"
		} else if r <= '1' && r >= '9' {
			return "Invalid Input\n"
		}
	}
	return "Invalid Input\n"
}

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("1454"))
}
