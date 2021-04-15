package main

import "fmt"

func InsertionSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return nil
	}
	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				swap(arr, j, j-1)
			} else {
				break
			}
		}
	}
	return arr
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	a := []int{2, 1, 3, 8, 5, 7, 6}
	res := InsertionSort(a)
	fmt.Println(res)
}
