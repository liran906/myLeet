// https://leetcode.cn/problems/last-stone-weight-ii/
// m

package dp

// 二维数组 dp
// 本题目标是将一堆石头分成两堆，使得两堆重量差最小。其实和 101 几乎一样。
//
//	转换思路：将其转化为 背包问题：在不超过总重一半的情况下，选出尽量重的一堆。
//	使用 0-1 背包求得最接近 sum/2 的子集和 x，另一堆就是 sum - x。
//	最小差值为：(sum - x) - x = sum - 2*x
func lastStoneWeightII_2d(stones []int) int {
	// 1. 计算总和
	sum := 0
	for i := range stones {
		sum += stones[i]
	}

	// 2. 问题转化为：从 stones 中选一些石头，总重量不超过 sum/2，且尽量接近 sum/2
	target := sum / 2

	// 3. dp[i][j] 表示前 i 个石头中，选出重量不超过 j 的组合，其最大重量是多少
	dp := make([][]int, len(stones))
	for i := range dp {
		dp[i] = make([]int, target+1)
	}

	// 4. 初始化第一个石头：当容量 j >= stones[0] 时，才可以选择它
	for j := stones[0]; j <= target; j++ {
		dp[0][j] = stones[0]
	}

	// 5. 状态转移
	for i := 1; i < len(stones); i++ {
		for j := 0; j <= target; j++ {
			if j < stones[i] {
				// 当前背包容量不足，不能选第 i 个石头，只能继承上一个状态
				dp[i][j] = dp[i-1][j]
			} else {
				// 两种选择：不选 或 选
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-stones[i]]+stones[i])
			}
		}
	}

	// 6. dp[n-1][target] 是最接近 sum/2 的一个组合和
	// 两堆石头重量分别是 x 和 sum - x，最终剩下的石头重量是两者的差
	return sum - dp[len(stones)-1][target]*2
}

// 1d array dp
func lastStoneWeightII(stones []int) int {
	// 求所有石头总重量
	sum := 0
	for i := range stones {
		sum += stones[i]
	}

	// 将问题转化为：尽可能将部分石头重量接近 sum/2
	target := sum / 2

	// 定义一维 dp 表：dp[j] 表示容量为 j 时，能达到的最大重量
	dp := make([]int, target+1)

	// 遍历每一块石头（物品）, 注意这里从 0 遍历，就不能初始化数组（因为会造成物品（stone[0]）的重复使用，导致结果错误❌）
	for i := range stones {
		// 倒序遍历背包容量，防止同一个物品被重复使用（0-1 背包的关键）
		for j := target; j >= 0; j-- {
			if j >= stones[i] {
				// 状态转移：选 or 不选当前石头
				dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
			}
		}
	}

	// sum - 2*dp[target] 即是两个子集重量差的最小值（最优结果）
	return sum - dp[target]*2
}
