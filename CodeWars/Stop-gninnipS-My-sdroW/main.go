package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(SpinWords(""))
}

func SpinWords(str string) string {
	var result string
	s := strings.Split(str, " ")
	for i := 0; i < len(s); i++ {
		if len(s[i]) >= 5 {
			s[i] = reverse(s[i])
		}
		result += s[i] + " "
	}

	return result
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
