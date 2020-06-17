package main

import (
	"fmt"
)

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	fmt.Println(twoSumGo([]int{3, 2, 4}, 6))
}

func twoSumLoop(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	var result []int

LOOP:
	for i1, v1 := range nums {
		for i2, v2 := range nums[i1+1:] {
			if v1+v2 == target {
				result = append(result, i1, i2+i1+1)
				break LOOP
			}
		}
	}
	return result
}

func twoSumHash(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	var result []int
	hash := make(map[int]int, len(nums))
	for i1, v1 := range nums {
		v2 := target - v1
		if i2, exist := hash[v2]; exist {
			result = append(result, i1, i2)
			break
		}
		hash[v1] = i1
	}
	return result
}

func twoSumGo(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	var result []int
	ch := make(chan []int)
	countCh := make(chan int)
	count := 0
	find := func(i, v, target int, nums []int) {
		for in, va := range nums {
			if va+v == target {
				result = append(result, i, i+in+1)
				ch <- result
				break
			}
		}
		countCh <- 1
	}
	for i, v := range nums {
		go find(i, v, target, nums[i+1:])
	}
	Loop:
	for {
		select {
		case result = <-ch:
			break Loop
		case <-countCh:
			count++
			if len(nums) == count {
				break Loop
			}
		}
	}
	return result
}
