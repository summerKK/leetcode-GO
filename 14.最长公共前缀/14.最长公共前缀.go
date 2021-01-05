package leetcode

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 找一个基准数据.
	base := strs[0]
	// 把数组的其他元素和它作对比.
	for i := 1; i < len(strs); i++ {
		// 对数组里面的每一个元素和基准作比较
		// 比如: summer(基准)  summary
		// 第一次比较 summer  summary  (-1)
		// 第二次比较 summe  summary  (-1)
		// 第三次比较 summ  summary  (0)
		for strings.Index(strs[i], base) != 0 {
			if len(base) == 0 {
				return ""
			}
			// 截掉最后一位
			base = base[:len(base)-1]
		}
	}

	return base
}
