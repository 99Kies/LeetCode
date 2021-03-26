//package main
//
//import "fmt"
//
//func createTargetArray(nums []int, index []int) []int {
//	//var result []int
//	//var temp []int
//	for i := range nums {
//		fmt.Println(nums[i], index[i])
//	}
//	return []int{0}
//}
//
//
//func main() {
//	a := []int{40,30,20,10}
//	c := [][]int{{1,2,3,4}}
//	fmt.Println(c)
//	b := []int{4,3,2,1}
//	createTargetArray(a, b)
//}
//
//
//import io
//import sys
//
//sys.stdout = io.TextIOWrapper(sys.stdout.buffer,encoding='utf-8')
//#str = input()
//#print(str)
//
//def sum(alists, index):
//sum = 0
//for alist in alists:
//for row in alist:
//sum = sum + row
//return sum
//
//a = [[12, 32, 9, 11, 34],
//[8,54, 76,23,7]]
//result = sum(a)
//print(result)

package main

//func sumlistbyindex(nums [][]int, index [][]int){
//	//var temp [][]int
//	row := len(nums)
//	col := len(nums[0])
//	flag := 0
//	irow := 0
//	//icol := 0
//	for flag < row*col{
//		//fmt.Printf("%d ", flag)
//		fmt.Printf("%d ", flag/col)
//		fmt.Printf("%d | ", flag%col)
//		if flag % col == col-1 {
//			fmt.Println()
//			irow++
//		}
//		//for i := 0 ;i<flag%col;i++{
//		//	fmt.Printf("%d ", flag)
//		//}
//		//if index[]
//		if index[irow][flag%col] == flag {
//
//			flag++
//		}else {
//			irow++
//		}
//	}
//
//
//	for i := range index{
//		for j:= range index[i]{
//			if index[i][j] == flag {
//				fmt.Printf("%d ", nums[i][(flag-1)%index[i][j]])
//				flag++
//			}
//			//fmt.Printf("%d ", index[i][j])
//			//fmt.Println(i)
//		}
//		flag = 1
//		fmt.Println()
//	}

//func main() {
//	//a := 0
//	//fmt.Scan(&a)
//	//fmt.Printf("%d\n", a)
//
//	//fmt.Printf("Hello World!\n");
//
//	//a := [][]int{{12, 32, 9, 11, 34}, {8,54, 76,23,7}}
//	//index := [][]int{{1, 2, 3, 4, 5}, {10, 9, 8, 7, 6}}
//	a := [][]int{{1,2,3,4,5},
//		{22,18,14,10,6},
//		{23,19,15,11,7},
//		{24,20,16,12,8},
//		{25,21,17,13,9}}
//	index := [][]int{{1,2,3,4,5},
//		{22,18,14,10,6},
//		{23,19,15,11,7},
//		{24,20,16,12,8},
//		{25,21,17,13,9}}
//	sumlistbyindex(a, index)
//}
