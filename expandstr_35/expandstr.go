package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// должен быть ровно один аргумент
	if len(os.Args) != 2 {
		return
	}

	s := os.Args[1]
	inWord := false
	printedWord := false

	for i := 0; i < len(s); i++ {
		// разделители слов
		if s[i] == ' ' || s[i] == '\t' {
			inWord = false
			continue
		}

		// начало нового слова (и это не первое слово)
		if !inWord && printedWord {
			z01.PrintRune(' ')
			z01.PrintRune(' ')
			z01.PrintRune(' ')
		}

		// печатаем символ слова
		z01.PrintRune(rune(s[i]))
		inWord = true
		printedWord = true
	}

	// если не было ни одного слова — ничего не печатаем
	if !printedWord {
		return
	}

	// newline в конце
	z01.PrintRune('\n')
}
