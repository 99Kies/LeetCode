package main

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			// 如果nums[0] < nums[mid] 则代表 左 侧是有序的列表
			if nums[0] <= target && target < nums[mid] {
				// 如果第一个值比target小，比mid点小则代表target在左侧
				r = mid - 1
			} else {
				// 反之target在右半边
				l = mid + 1
			}
		} else {
			// 如果nums[0] > nums[mid] 则代表 右 侧是有序的列表
			if nums[mid] < target && target <= nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		// if ((nums[0] > target)  (num) )
	}
	return -1
}

// @lc code=end

func main() {

}
