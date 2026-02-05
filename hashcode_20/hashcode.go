package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

func HashCode(dec string) string {
	result := ""
	size := len(dec)

	for _, r := range dec {
		hash := (int(r) + size) % 127

		if hash < 32 {
			hash += 33
		}
		result += string(rune(hash))
	}
	return result
}
