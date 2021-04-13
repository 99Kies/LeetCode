package main

import "fmt"

type SubrectangleQueries struct {
	tangle [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{
		rectangle,
	}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			this.tangle[i][j] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.tangle[row][col]
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */

func main() {
	rectangle := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	obj := Constructor(rectangle)
	fmt.Println(obj.tangle)
	obj.UpdateSubrectangle(0, 1, 1, 2, 2)
	param_2 := obj.GetValue(0, 0)
	fmt.Println(obj.tangle)
	fmt.Println(param_2)
	param_2 = obj.GetValue(1, 2)
	fmt.Println(param_2)
}
