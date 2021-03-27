package main

func numberOfMatches(n int) int {
	step := n
	sum := 0
	for step > 1 {
		sum = step/2 + sum
		step = step - step/2
	}
	return sum
}

func main() {
	numberOfMatches(7)
}
