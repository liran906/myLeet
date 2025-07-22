// https://leetcode.cn/problems/integer-break/description/
// m

package dp

func integerBreak(n int) int {
	// 当 n < 4 时，按照题目要求至少拆分成两个正整数
	// 特殊处理：2 -> 1（1+1）、3 -> 2（1+2）
	if n < 4 {
		return n - 1
	}

	// 创建 dp 数组，dp[i] 表示将正整数 i 拆分后的最大乘积
	dp := make([]int, n+1)

	for i := 2; i <= n; i++ {
		// 对于小于 5 的数，初始化为本身值
		// 这样 dp[2]=2，dp[3]=3，dp[4]=4，可以避免在乘法时被劣化
		if i < 5 {
			dp[i] = i
			continue
		}

		// 枚举所有可能的拆分 j + (i - j)，其中 j 从 2 到 i/2
		// 注意：由于乘法交换律，对称一半就够了
		for j := 2; j <= i/2; j++ {
			dp[i] = max(dp[i], dp[j]*dp[i-j])
			// 这里如果用正常写法
			// 也就是 dp[2] dp[3]我前面不人为赋值为 2 和 3，而是 1 和 2
			// 那么这里写作：
			// dp[i] = max(dp[i], max(dp[j], j) * max(dp[i-j], i-j))
		}
	}

	// 题目要求的是将 n 拆分后最大乘积
	return dp[n]
}
