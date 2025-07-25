// https://leetcode.cn/problems/partition-equal-subset-sum/description/
// m

package dp

import "sort"

// 错误方法：想用双指针，反例[1,1,2,2]
func canPartition_wrong(nums []int) bool {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	sum := 0
	for l <= r {
		if sum < 0 {
			sum += nums[l]
			l++
		} else {
			sum -= nums[r]
			r--
		}
	}
	return sum == 0
}

// https://programmercarl.com/0416.%E5%88%86%E5%89%B2%E7%AD%89%E5%92%8C%E5%AD%90%E9%9B%86.html
// 二维数组dp
func canPartition_2d(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	var dp = make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	for j := nums[0]; j <= target; j++ {
		dp[0][j] = nums[0]
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j <= target; j++ {
			if j < nums[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
			}
			if dp[i][j] == target {
				return true
			}
		}
	}
	return dp[len(nums)-1][target] == target
}

// 一维数组db 比 2d 快不少，因为这题 j 的大小是跟 sum 挂钩而不是len（nums），但其实两个的复杂度是一样的
func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	var dp = make([]int, target+1)
	for i := 1; i < len(nums); i++ {
		for j := target; j >= 0; j-- {
			if j >= nums[i] {
				dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
			}
			if dp[j] == target {
				return true
			}
		}
	}
	return dp[target] == target
}
