package main

import "fmt"

// XOR 异或

func decode(encoded []int, first int) []int {
	length := len(encoded)
	res := make([]int, length+1)
	res[0] = first
	for i := 0; i < length; i++ {
		res[i+1] = res[i] ^ encoded[i]
	}
	return res
}

func main() {
	encoded := []int{6, 2, 7, 3}
	first := 4
	fmt.Println(decode(encoded, first))
}
