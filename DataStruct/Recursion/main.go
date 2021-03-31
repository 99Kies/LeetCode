package main

import "fmt"

func Feibonaqi(n int) int {

	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Feibonaqi(n-1) + Feibonaqi(n-2)
}

func main() {
	a := Feibonaqi(50000)
	fmt.Println(a)
}
