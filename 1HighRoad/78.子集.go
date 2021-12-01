package main

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
var res = [][]int{}

func subsets(nums []int) [][]int {
	// 设置一个返回的存储
	res = make([][]int, 0)
	n := len(nums)
	backtrace(nums, []int{}, 0, n)
	return res
	// var dfs func(i int, list []int)
	// dfs = func(i int, list []int) {
	// 	if i == len(nums) {
	// 		tmp := make([]int, len(list))
	// 		copy(tmp, list) // 只复制内容，不复制指针。
	// 		res = append(res, tmp)
	// 		return
	// 	}
	// 	list = append(list, nums[i])
	// 	dfs(i+1, list)
	// 	list = list[:len(list)-1]
	// 	dfs(i+1, list)
	// }
	// dfs(0, []int{})

	// return res
}

func backtrace(nums, path []int, start int, n int) {
	res = append(res, path)
	for i := start; i < n; i++ {

		tmp := make([]int, len(path))
		copy(tmp, path) // 只复制内容，不复制指针。
		path = append(tmp, nums[i])
		backtrace(nums, path, i+1, n)
		path = path[:len(path)-1]

	}

}

// @lc code=end

func main() {

}
