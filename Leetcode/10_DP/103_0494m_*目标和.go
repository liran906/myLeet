// https://leetcode.cn/problems/target-sum/
// m

package dp

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	// 把 sum 分为加正号的 right 和加负号的 left，那么：
	// right + left = sum
	// right - left = target
	// 联立可得 right = (sum + target) / 2
	// right 实际就是 dp 中的背包上限，本题中用 upbound
	if (sum+target)%2 == 1 {
		// 此情况无解
		return 0
	}

	if abs(target) > sum {
		// 如果 target 本身超过了 sum，也无解
		return 0
	}
	// 开始 dp
	// 这里写一维方法，原理上二维更好理解，可以看卡码笔记
	upbound := (sum + target) / 2
	dp := make([]int, upbound+1)
	dp[0] = 1 // 初始化，重量为 0 的时候，啥都不放
	// dp[nums[0]] = 1 // 初始化，重量为第一个数时有一种放法

	for i := 1; i < len(nums); i++ {
		for j := upbound; j >= 0; j-- {
			// 减去当前物品容量为负，那就不能放当前物品
			// 在一维数组 dp 写法就是 dp[j] = dp[j]
			// 所以可以省略不写了
			if j >= nums[i] {
				// 考虑两个情况
				// 1 放物品 i：和减去物品 i 重量的放置方法数量一致
				// 2 不放物品 i：和当前重量下不考虑物品 i 的放置方法数量一致
				// 最终的结果就是情况 1 +情况 2
				dp[j] = dp[j] + dp[j-nums[i]]
			}
		}
	}
	return dp[upbound]
}

// 辅助函数：计算绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
