package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5) == 2)
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2), searchInsert([]int{1, 3, 5, 6}, 2) == 1)
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7), searchInsert([]int{1, 3, 5, 6}, 7) == 4)
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0) == 0)
}

/**
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2
示例 2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
示例 4:

输入: [1,3,5,6], 0
输出: 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-insert-position
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	idx := 0
	for l <= r {
		idx = (l + r) / 2
		if target == nums[idx] {
			return idx
		} else if target > nums[idx] {
			l = idx + 1
		} else if target < nums[idx] {
			r = idx - 1
		}
	}
	if nums[idx] < target {
		return idx + 1
	}
	return idx
}
