// https://leetcode.cn/problems/maximum-subarray/description/
// m

package greedy

// 贪心
func maxSubArray(nums []int) int {
	var res, cur = nums[0], 0 // 结果先设为第一个个数
	for i := range nums {     // 遍历数组
		cur += nums[i] // 加上当前的数字
		if cur > res { // 如果比结果大
			res = cur // 那就更新结果
		}
		if cur < 0 { // 如果当前和小于零
			cur = 0 // 重置当前和为 0
		}
	}
	return res
}
