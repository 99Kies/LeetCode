package main

import "fmt"

func minPartitions(n string) int {
	max := 0
	for _, i := range n {
		if int(i-'0') > max {
			max = int(i - '0')
		}
	}
	return max
}

func main() {
	num := "32"
	res := minPartitions(num)
	fmt.Println(res)
}
