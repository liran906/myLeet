// https://leetcode.cn/problems/house-robber/
package dp

// 常规动态规划，这里是一维的
func rob_1d(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

// 一维的可以压缩到常量纬，就跟二维压缩到一维一样的道理
func rob_(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	prev2 := nums[0]
	prev1 := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		cur := max(prev2+nums[i], prev1)
		prev2 = prev1
		prev1 = cur
	}
	return prev1
}
