package main

import "fmt"

//type Node struct {
//	Start
//}
//
//
//type QueenStack struct {
//	element [][]int
//}

func main() {
	test_num := 0
	fmt.Scanln(&test_num)
	all_result := make([][]int, test_num)
	n := 0
	for i := 0; i < test_num; i++ {
		fmt.Scanln(&n)
		L := 0
		R := 0
		queenNums := make([][]int, n)
		result := make([]int, n)
		max := 0
		for j := 0; j < n; j++ {
			fmt.Scanln(&L, &R)
			queenNums[j] = append(queenNums[j], L)
			queenNums[j] = append(queenNums[j], R)
		}
		for j := 0; j < n; j++ {
			if queenNums[j][0] == 1 && j == 0 {
				result[0] = 1
				max = result[j]
			}
			if j > 0 {
				if queenNums[j][1] > queenNums[j-1][0] {
					result[j] = max + 1
					max = result[j]
				} else if queenNums[j][1] < queenNums[j-1][0] {
					result[j] = 0
					//}else{
					//	result[j] = max  + 1
					//	max = result[j]
				}
			}
		}

		//fmt.Println(result)
		all_result[i] = result

		//fmt.Scanln(&)
		//?	fmt.Println(i)
	}
	//fmt.Println(all_result)
	for i := 0; i < test_num; i++ {
		length := len(all_result[i])
		//fmt.Printf(all_result[i])
		for j := 0; j < length; j++ {
			if j < length-1 {
				fmt.Printf("%d ", all_result[i][j])
			} else {
				fmt.Printf("%d\n", all_result[i][j])
			}
		}
	}
}
