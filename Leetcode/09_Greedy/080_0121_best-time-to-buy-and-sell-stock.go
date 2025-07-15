// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
// e

package greedy

func maxProfit_i(prices []int) int {
	start, profit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		profit = max(profit, prices[i]-start)
		start = min(start, prices[i])
	}
	return profit
}
