package main

import "fmt"

func numberOfSteps(num int) int {
	k := num
	cnt := 0
	for k > 0 {
		if k%2 == 1 {
			k = k - 1
		} else {
			k = k / 2
		}
		cnt++
	}
	return cnt
}

func main() {
	num := 14
	anw := numberOfSteps(num)
	fmt.Println(anw)
}
