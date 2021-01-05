package leetcode

// 每天都存在交易.只要保证后一天比前一天价格高就交易
func maxProfit(prices []int) int {
	l := len(prices)
	profit := 0
	for i := 1; i < l; i++ {
		if d := prices[i] - prices[i-1]; d > 0 {
			profit += d
		}
	}

	return profit
}
