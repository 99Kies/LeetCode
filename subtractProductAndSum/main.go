package main

import "fmt"

func subtractProductAndSum(n int) int {
	sum := 0
	product := 1
	for n > 0 {
		num := n % 10
		sum = sum + num
		product = product * num

		n = n / 10
	}

	return product - sum

}

func main() {
	a := subtractProductAndSum(4421)
	fmt.Println(a)
}
