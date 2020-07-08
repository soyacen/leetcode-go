package main

import (
	"fmt"
)

/**

给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func main() {
	fmt.Println(reverse(123), reverse(123) == 321)
	fmt.Println(reverse(-123), reverse(-123) == -321)
	fmt.Println(reverse(120), reverse(120) == 21)

}

func reverse(x int) int {
	MAX := 1<<31 - 1
	MIN := -1 << 31
	res := 0
	for x != 0 {
		num := x % 10
		x = x / 10
		if MAX/10 < res || (MAX/10 == res && MAX%10 < num) {
			return 0
		}
		if MIN/10 > res || (MIN/10 == res && MIN%10 > num) {
			return 0
		}
		res = res*10 + num
	}
	return res
}
