package main

import "fmt"

/**
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zigzag-conversion
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(convert("LEETCODEISHIRING", 3))
	fmt.Println(convert("LEETCODEISHIRING", 3) == "LCIRETOESIIGEDHN")
	fmt.Println(convert("LEETCODEISHIRING", 4) == "LDREOEIIECIHNTSG")
	fmt.Println(convert("ABCDE", 4))

}

func convert1(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	seq := []rune(s)
	leng := len(s)
	rowLen := (leng + numRows*2 - 2) / (numRows*2 - 2) * (numRows - 1)
	matrix := make([][]rune, numRows)
	for i := range matrix {
		matrix[i] = make([]rune, rowLen)
	}
	flag := 0
	row, col := 0, -1
	for i, c := range seq {
		flag = i % (numRows*2 - 2)
		if flag == 0 || flag >= numRows {
			col++
		}
		if flag < numRows {
			row = flag
		} else {
			row--
		}
		matrix[row][col] = c
	}
	fmt.Println(matrix)

	res := make([]rune, 0, leng)
	for _, row := range matrix {
		for _, c := range row {
			if c == 0 {
				continue
			}
			res = append(res, c)
		}
	}

	return string(res)
}

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	seq := []rune(s)
	rowCount := Min(numRows, len(seq))
	rows := make([][]rune, rowCount)
	for i :=range rows {
		rows[i] = make([]rune, 0)
	}

	rowIndex := 0
	isGoDown:= false
	for _, c := range seq {
		rows[rowIndex] = append(rows[rowIndex], c)
		if rowIndex == 0 || rowIndex == rowCount -1 {
			isGoDown = !isGoDown
		}
		if isGoDown {
			rowIndex ++
		} else {
			rowIndex --
		}
	}
	
	result := make([]rune,0, len(seq))
	for _, row := range rows {
		result = append(result, row...)
	}
	return string(result)
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
