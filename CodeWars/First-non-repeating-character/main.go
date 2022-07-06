package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(FirstNonRepeating("sdfhushfiadjdshisghfiuew"))
}

func FirstNonRepeating(str string) string {

	iter := make([]int, len(str))
	for i := 0; i < len(str); i++ {

		for j := 0; j < len(str); j++ {
			if (i != j) && (strings.ToLower(string(str[i])) == strings.ToLower(string(str[j]))) {
				iter[i]++
			}
		}
		if iter[i] == 0 {
			return string(str[i])
		}
	}
	return ""
}
