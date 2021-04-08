package main

import "fmt"

func waysToChange(n int) int {
	mod := 10 ^ 9 + 7
	coin := []int{25, 10, 5, 1}
	f := make([]int, n+1)
	f[0] = 1
	//all_sum_coin := 0
	//for i:=0; all_sum_coin == n ;i++{
	//
	//}
	//
	for _, j := range coin {
		for i := 1; i < n+1; i++ {
			if i >= j {
				f[i] = (f[i] + f[i-j]) % mod
			}
		}
	}
	return f[n]
}

func coinChange(A []int, M int) int {
	f := make([]int, M+1)
	n := len(A)

	f[0] = 0

	for i := 1; i <= M; i++ {
		f[i] = 999
		for j := 0; j < n; j++ {
			if i >= A[j] && f[i-A[j]] != 999 {
				f[i] = min(f[i-A[j]], f[i])
			}
		}
	}
	if f[M] == 999 {
		f[M] = -1
	}
	return f[M]
}

func min(x, y int) int {
	if x >= y {
		return y
	}
	return x
}

func main() {
	a := waysToChange(5)
	fmt.Println(a)
}
