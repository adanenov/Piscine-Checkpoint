package main

import (
	"fmt"
)

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}

func FromTo(from int, to int) string {
	// проверка диапазона
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	result := ""

	// если from == to
	if from == to {
		if from < 10 {
			return "0" + itoa(from) + "\n"
		}
		return itoa(from) + "\n"
	}

	// движение вперёд
	if from < to {
		for i := from; i <= to; i++ {
			if i < 10 {
				result += "0" + itoa(i)
			} else {
				result += itoa(i)
			}
			if i != to {
				result += ", "
			}
		}
		result += "\n"
		return result
	}

	// движение назад
	for i := from; i >= to; i-- {
		if i < 10 {
			result += "0" + itoa(i)
		} else {
			result += itoa(i)
		}
		if i != to {
			result += ", "
		}
	}
	result += "\n"

	return result
}

// простая конвертация int → string
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	res := ""
	for n > 0 {
		res = string(rune('0'+n%10)) + res
		n /= 10
	}
	return res
}
