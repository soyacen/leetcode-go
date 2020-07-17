package main

import "fmt"

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}

/*
给你两个二进制字符串，返回它们的和（用二进制表示）。

输入为 非空 字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"

提示：
每个字符串仅由字符 '0' 或 '1' 组成。
1 <= a.length, b.length <= 10^4
字符串如果不是 "0" ，就都不含前导零。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-binary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func addBinary(a string, b string) string {
	ab := []byte(a)
	lenA := len(ab)
	bb := []byte(b)
	lenB := len(bb)
	length := lenA + 1
	if lenB > lenA {
		length = lenB + 1
	}
	result := make([]byte, length)
	var carry byte = 0
	i := 0
	for {
		if lenA <= i && lenB <= i {
			break
		}
		var x byte = 0
		if lenA > i {
			x = ab[lenA-1-i] - '0'
		}
		var y byte = 0
		if lenB > i {
			y = bb[lenB-1-i] - '0'
		}
		sum := x + y + carry
		result[length-i-1] = sum%2 + '0'
		carry = sum / 2
		i++
	}
	if carry != 0 {
		result[0] = carry + '0'
		return string(result)
	}

	return string(result[1:])
}
