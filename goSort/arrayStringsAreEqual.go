package main

import "fmt"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var a string
	var b string
	for i := range word1 {
		a = a + word1[i]
	}
	for j := range word2 {
		b = b + word2[j]
	}
	if a == b {
		return true
	}
	return false
}

func main() {
	a := []string{"a", "bc"}
	b := []string{"ab", "c"}
	c := []string{"ab", "bc"}
	fmt.Println(arrayStringsAreEqual(a, b))
	fmt.Println(arrayStringsAreEqual(a, c))
}
