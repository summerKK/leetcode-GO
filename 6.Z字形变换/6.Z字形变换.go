package leetcode

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	// 字符串装换为数组
	chars := []rune(s)
	// 一个周期
	period := 2*numRows - 2
	result := make([]string, numRows)
	for i, char := range chars {
		index := i % period
		if index < numRows {
			result[index] += string(char)
		} else {
			result[period-index] += string(char)
		}
	}

	return strings.Join(result, "")
}
