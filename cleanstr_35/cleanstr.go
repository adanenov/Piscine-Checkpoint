package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}
	s := os.Args[1]
	inWord := false
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' {
			if inWord {
				z01.PrintRune(' ')
				inWord = false
			}
		} else {
			z01.PrintRune(rune(s[i]))
			inWord = true
		}
	}
	z01.PrintRune('\n')
}
