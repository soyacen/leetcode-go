package main

import "fmt"

func main() {
	fmt.Println(climbStairs1(2))
	fmt.Println(climbStairs1(3))
	fmt.Println(climbStairs1(44))
	fmt.Println(climbStairs2(2))
	fmt.Println(climbStairs2(3))
	fmt.Println(climbStairs2(44))
}

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func climbStairs1(n int) int {
	if n <= 1 {
		return 1
	}
	var count int
	count += climbStairs1(n - 1)
	count += climbStairs1(n - 2)

	return count
}

// f(x) = f(x-1) + f(x-2)
// f(1) = 1
// f(2) = 2
func climbStairs2(n int) int {
	r, x, y := 0, 0, 0
	for i := 1; i <= n; i++ {
		if i == 1 {
			r = 1
			x = 1
			continue
		}
		if i == 2 {
			r = 2
			y = 2
			continue
		}
		r = x + y
		x = y
		y = r
	}
	return r
}
