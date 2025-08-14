//https://leetcode.cn/problems/jump-game-ii/description/
// m

package greedy

func jump(nums []int) (steps int) {
	var curMaxDis int // 当前能跳到的最远距离（当前层的最远边界）
	var nextDis int   // 下一层能跳到的最远距离

	for i := 0; i < len(nums); i++ {
		// 如果当前能到达的最远距离已经到达或超过终点，直接跳出
		if curMaxDis >= len(nums)-1 {
			break
		}

		// 更新从当前位置出发所能到达的最远位置
		nextDis = max(nextDis, i+nums[i])

		// 如果走到当前层的边界，准备跳到下一层
		if i == curMaxDis {
			curMaxDis = nextDis // 切换到下一层的边界
			steps++             // 计数跳跃次数
		}
	}
	return
}

// 动态规划算法 On^2 没有贪心好
func jump_dp(nums []int) (steps int) {
	n := len(nums)
	dp := make([]int, n)
	// 初始化
	for i := 1; i < n; i++ {
		dp[i] = n
	}

	// dp
	for i := 0; i < n; i++ {
		maxLeap := min(i+nums[i], n-1) // 本次最大跳跃位置的索引
		for j := i + 1; j <= maxLeap; j++ {
			dp[j] = min(dp[j], dp[i]+1)
		}
	}
	return dp[n-1]
}
