// https://leetcode.cn/problems/min-cost-climbing-stairs/
// e

package dp

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	// dp[i]表示爬到第 i 层的最少花费
	// dp[0] = 0 到达下标 0 的台阶花费为 0
	// dp[1] = 0 到达下标 1 的台阶花费也为 0

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}
