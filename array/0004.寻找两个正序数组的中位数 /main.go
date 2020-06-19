package main

import (
	"fmt"
	"math"
)

/**
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。

请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。



示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{1}, []int{}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))

	fmt.Println("=============")
	fmt.Println(findMedianSortedArraysLog([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArraysLog([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArraysLog([]int{1}, []int{}))
	fmt.Println(findMedianSortedArraysLog([]int{}, []int{1}))
}

func findMedianSortedArraysLog(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	left, right := 0, m
	median1, median2 := 0, 0
	for left <= right {
		i := (left + right) / 2
		j := (m+n+1)/2 - i
		nums_im1 := math.MinInt32
		if i != 0 {
			nums_im1 = nums1[i-1]
		}
		nums_i := math.MaxInt32
		if i != m {
			nums_i = nums1[i]
		}
		nums_jm1 := math.MinInt32
		if j != 0 {
			nums_jm1 = nums2[j-1]
		}
		nums_j := math.MaxInt32
		if j != n {
			nums_j = nums2[j]
		}
		if nums_im1 <= nums_j {
			median1 = max(nums_im1, nums_jm1)
			median2 = min(nums_i, nums_j)
			left = i + 1
		} else {
			right = i - 1
		}
	}
	if (m+n)%2 == 0 {
		return float64(median1+median2) / 2.0
	}
	return float64(median1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length1 := len(nums1)
	length2 := len(nums2)
	index1, index2 := 0, 0
	totalLen := length1 + length2
	count := 0

	odd := false
	if totalLen%2 != 0 {
		odd = true
	}
	mid := totalLen / 2

	n := -1
	for {
		curr := -1
		if index1 < length1 && index2 < length2 {
			if nums1[index1] < nums2[index2] {
				curr = nums1[index1]
				index1++
			} else {
				curr = nums2[index2]
				index2++
			}
		} else if index1 < length1 {
			curr = nums1[index1]
			index1++
		} else if index2 < length2 {
			curr = nums2[index2]
			index2++
		} else {
			break
		}
		if count == mid && odd {
			return float64(curr)
		}
		if count == mid-1 && !odd {
			n = curr
		}
		if count == mid && !odd {
			return (float64(n) + float64(curr)) / 2
		}
		count++
	}
	return 0
}
