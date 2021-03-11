package main

import (
	"fmt"
	"strings"
)

func countConsistentStrings(allowed string, words []string) int {

	for j := range words {
		b := strings.Contains(allowed, words[j])
		fmt.Println(b)
	}
	return 0
}

func main() {
	a := []string{"a", "ab", "abc"}
	countConsistentStrings("abcd", a)
}
