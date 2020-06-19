package main

import "fmt"

/**
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
通过次数295,784提交次数961,335

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("ac"))
	//fmt.Println(checkPalindromeInRuneArray([]rune("c"), 0, 0))
	//fmt.Println(checkPalindromeInRuneArray([]rune("ac"), 0, 1))
	//fmt.Println(checkPalindromeInRuneArray([]rune("cbbd"), 0, 3))
	//fmt.Println(checkPalindromeInRuneArray([]rune("dbbd"), 0, 3))
	//fmt.Println(checkPalindromeInRuneArray([]rune("db1d"), 0, 3))
	//fmt.Println(checkPalindromeInRuneArray([]rune("dbcbd"), 0, 4))

}

func longestPalindrome(s string) string {
	return string(longestPalindromeArray([]rune(s)))
}

func longestPalindromeArray(runes []rune) []rune {
	for i := 0; i < len(runes); i++ {
		length := len(runes) - i
		for j := 0; j < i+1; j++ {
			if checkPalindromeInRuneArray(runes, j, length+j-1) {
				return runes[j : j+length]
			}
		}
	}
	return nil
}

func checkPalindromeInRuneArray(runes []rune, i int, j int) bool {
	if i >= j || len(runes) == 0 {
		return true
	}
	if runes[i] == runes[j] {
		return checkPalindromeInRuneArray(runes, i+1, j-1)
	} else {
		return false
	}
}
