package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubArray3([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), maxSubArray3([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) == 6)
}

/**
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 动态规划
func maxSubArray1(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// 贪心
func maxSubArray2(nums []int) int {
	max := -1 >> 31
	sum := 0
	for _, n := range nums {
		sum += n
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}

// 分治
func maxSubArray3(nums []int) int {
	return findMax(nums, 0, len(nums)-1)
}

func findMax(nums []int, l int, r int) int {
	if l == r {
		return nums[0]
	}
	// 中心点
	mid := (l + r) / 2
	// 左边
	maxLeft := findMax(nums, l, mid)
	// 右边
	maxRight := findMax(nums, mid+1, r)

	// 跨中心
	maxMid := findMaxMid(nums, l, mid, r)
	fmt.Println(maxMid, l, mid, r)
	result := maxLeft
	if result < maxRight {
		result = maxRight
	}
	if result < maxMid {
		result = maxMid
	}
	return result
}

func findMaxMid(nums []int, l int, mid int, r int) int {
	leftMax := nums[mid]
	sum := nums[mid]
	for i := mid - 1; i >= l; i-- {
		sum += nums[i]
		if sum > leftMax {
			leftMax = sum
		}
	}

	rightMax := nums[mid+1]
	sum = nums[mid+1]
	for i := mid + 2; i <= r; i++ {
		sum += nums[i]
		if sum > rightMax {
			rightMax = sum
		}
	}
	return leftMax + rightMax
}
