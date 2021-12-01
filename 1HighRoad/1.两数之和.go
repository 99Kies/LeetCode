package main

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	hashTable[nums[0]] = 0
	for i := 1; i < len(nums); i++ {
		if _, ok := hashTable[target-nums[i]]; ok {
			return []int{hashTable[target-nums[i]], i}
		}
		hashTable[nums[i]] = i
	}
	return []int{}
}

// @lc code=end

func main() {

}
