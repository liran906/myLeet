// https://leetcode.cn/problems/perfect-squares/description/

package dp

import "math"

func numSquares_(n int) int {
	// 计算不大于 n 的最大平方根（因为只需要用到 i^2 ≤ n）
	root := int(math.Sqrt(float64(n)))

	// dp[i] 表示：凑出 i 所需的最少完全平方数数量
	dp := make([]int, n+1)

	// 初始化：dp[0] = 0，其余设为无穷大（代表不可达）
	for i := range dp {
		if i > 0 {
			dp[i] = math.MaxInt32
		}
	}

	// 遍历每一个平方数 i*i
	for i := 1; i <= root; i++ {
		sq := i * i
		// 尝试用当前平方数更新所有 dp[j]
		for j := sq; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-sq]+1) // 关键点在这里：不是 dp[sq]，而是 +1
		}
	}

	return dp[n]
}

// 还可以这样，不开根
func numSquares(n int) int {
	//定义
	dp := make([]int, n+1)
	// 初始化
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	// 遍历物品
	for i := 1; i*i <= n; i++ {
		// 遍历背包
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}

	return dp[n]
}
