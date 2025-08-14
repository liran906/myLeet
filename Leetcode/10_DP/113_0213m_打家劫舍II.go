// https://leetcode.cn/problems/house-robber-ii/
package dp

// https://programmercarl.com/0213.%E6%89%93%E5%AE%B6%E5%8A%AB%E8%88%8DII.html
// 分两个情况考虑：不考虑头和不考虑尾，两个情况再取最大值
func robii(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	var robrange func(int, int) int
	robrange = func(start, end int) int {
		if start == end {
			return nums[start]
		}

		prev2 := nums[start]
		prev1 := max(nums[start], nums[start+1])

		for i := start + 2; i <= end; i++ {
			cur := max(prev2+nums[i], prev1)
			prev2 = prev1
			prev1 = cur
		}
		return prev1
	}

	res1 := robrange(0, n-2)
	res2 := robrange(1, n-1)
	return max(res1, res2)
}

// 不压缩的写法
// rob 计算环形房子的最大偷窃金额
// nums[i] 表示第 i 个房子的金额
func rob_dp_(nums []int) int {
	n := len(nums)
	if n == 0 {
		// 没有房子
		return 0
	}
	if n == 1 {
		// 只有一个房子，直接偷它
		return nums[0]
	}

	// robRange 计算 nums[start:end]（包含 start 和 end）这一段的最大偷窃金额
	var robRange func(int, int) int
	robRange = func(start, end int) int {
		length := end - start + 1
		if length == 1 {
			// 区间只有一个房子，直接偷它
			return nums[start]
		}

		// dp[i] 表示从 nums[start] 到 nums[start+i] 这一段的最大偷窃金额
		dp := make([]int, length)

		// 初始化前两个位置
		dp[0] = nums[start]                     // 偷第一个房子
		dp[1] = max(nums[start], nums[start+1]) // 偷第一个或第二个，取较大值

		// 从第三个房子（i=2）开始递推
		// 这里 i 是相对 dp 的下标（相对 start 偏移量）
		// 所以 i < length（而不是 <=），因为 dp 的最后一个索引是 length-1
		for i := 2; i < length; i++ {
			// 两种选择：
			// 1. 不偷当前房子：dp[i-1]
			// 2. 偷当前房子：dp[i-2] + 当前房子的金额
			dp[i] = max(dp[i-1], dp[i-2]+nums[start+i])
		}

		// 返回最后一个房子（end）对应的最大收益
		return dp[length-1]
	}

	// 环形房子：不能同时偷第一个和最后一个
	// 情况1：不偷最后一间（0 到 n-2）
	// 情况2：不偷第一间（1 到 n-1）
	// 两种情况取最大值
	return max(robRange(0, n-2), robRange(1, n-1))
}
