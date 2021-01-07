package leetcode

/**
根据题意加一，没错就是加一这很重要，因为它是只加一的所以有可能的情况就只有两种：

除 99 之外的数字加一；
数字 99。
加一得十进一位个位数为 00 加法运算如不出现进位就运算结束了且进位只会是一。

所以只需要判断有没有进位并模拟出它的进位方式，如十位数加 11 个位数置为 00，如此循环直到判断没有再进位就退出循环返回结果。

然后还有一些特殊情况就是当出现 9999、999999 之类的数字时，循环到最后也需要进位，出现这种情况时需要手动将它进一位

作者：yhhzw
链接：https://leetcode-cn.com/problems/plus-one/solution/java-shu-xue-jie-ti-by-yhhzw/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
// 以 599 为例
// 第一次循环(逆序) digits[l]++ = 10,digits[l] % 10 = 0 59[0]
// 第二次循环(逆序) digits[l]++ = 10,digits[l] % 10 = 0 5[0][0]
// 第三次循环(逆序) digits[l]++ = 6,digits[l] % 10 = 6 [6][0][0] 直接return

// 以 99 为例
// 第一次循环(逆序) digits[l]++ = 10,digits[l] % 10 = 0 9[0]
// 第二次循环(逆序) digits[l]++ = 10,digits[l] % 10 = 0 [0][0] 循环结束
// 退出循环
// 首部添加1 digits[0] = 1 [1][0][0]
func plusOne(digits []int) []int {
	for l := len(digits) - 1; l >= 0; l-- {
		digits[l]++
		// 1 % 10 = 1  10 % 10 = 0
		digits[l] = digits[l] % 10
		if digits[l] != 0 {
			return digits
		}
	}
	digits = append([]int{1}, digits...)

	return digits
}
