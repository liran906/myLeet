// https://leetcode.cn/problems/coin-change-ii/
// m

package dp

func change(amount int, coins []int) int {
	// 定义 dp[j] 表示凑成金额 j 的组合数
	// 这里选择把二维 dp 压缩为一维，所以只要初始化 dp[0]
	// 如果用二维 dp，那么整个第一行第一列都要初始化
	// 第一行全 1，第一列全 1
	var dp = make([]int, amount+1)

	// 初始化：凑成金额 0 的组合有一种 —— 什么也不选
	dp[0] = 1

	// 遍历每一种硬币（物品），外层循环为物品，确保组合不重复
	for i := 0; i < len(coins); i++ {
		// 内层倒序遍历背包容量，从小到大遍历是为了允许“完全”使用当前硬币（可重复）
		for j := coins[i]; j <= amount; j++ {
			// 状态转移方程：
			// dp[j] += dp[j - coins[i]]
			// 表示：在已有 dp[j - coins[i]] 的组合数基础上，加上当前硬币后的组合
			dp[j] = dp[j] + dp[j-coins[i]]
		}
	}

	// 最终结果：凑成金额 amount 的组合数
	return dp[amount]
}
