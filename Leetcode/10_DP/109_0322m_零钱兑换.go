// https://leetcode.cn/problems/coin-change/

package dp

import "math"

// 本题求钱币最小个数，那么钱币有顺序和没有顺序都可以，都不影响钱币的最小个数。
// 所以本题并不强调集合是组合还是排列。

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		if i == 0 {
			dp[i] = 0
		} else {
			dp[i] = math.MaxInt32
		}

	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}
	if dp[amount] == math.MaxInt32 {
		dp[amount] = -1
	}
	return dp[amount]
}
