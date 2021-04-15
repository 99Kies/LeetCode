package main

import "fmt"

func minOperations(n int) int {
	return n * n / 4
}

func main() {
	a := 3
	res := minOperations(a)
	fmt.Println(res)
}
