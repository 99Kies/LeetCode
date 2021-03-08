package main

import "fmt"

func xorOperation(n int, start int) int {
	var result int
	for i := 0; i < n; i++ {
		result = (start + i*2) ^ result
	}
	return result
}

func main() {

	fmt.Println(xorOperation(5, 0))

}
