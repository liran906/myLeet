// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
// m

package greedy

func maxProfit(prices []int) (profit int) {
	// 其实不会越界，所以这个条件不写也行
	// if len(prices) <= 1 {
	//     return
	// }
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return
}
