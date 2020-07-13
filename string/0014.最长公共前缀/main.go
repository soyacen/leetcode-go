package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}), longestCommonPrefix([]string{"flower","flow","flight"}) == "fl")
	fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}), longestCommonPrefix([]string{"dog","racecar","car"}) == "")
}

/**
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var result []byte
	idx := 0
	for {
		var char byte
		for i, str := range strs {
			if str == "" {
				return string(result)
			}
			if idx >=len(str) {
				return string(result)
			}
			if i == 0 {
				char = str[idx]
				continue
			}
			if char == str[idx] {
				continue
			} else {
				return string(result)
			}
		}
		result = append(result, char)
		char = ' '
		idx ++
	}
	return string(result)
}