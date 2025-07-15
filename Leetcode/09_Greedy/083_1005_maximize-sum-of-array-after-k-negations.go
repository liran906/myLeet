// https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations/
// m

package greedy

import "sort"

func largestSumAfterKNegations(nums []int, k int) (res int) {
	// 先排序
	sort.Ints(nums)

	// 翻转负数
	for i := 0; i < len(nums) && k > 0; i++ {
		if nums[i] >= 0 {
			break
		}
		nums[i] = -nums[i]
		k--
	}

	// 如果剩余的 k 为奇数，重新排序，翻转最小的
	if k%2 == 1 {
		sort.Ints(nums)
		nums[0] = -nums[0]
	}

	// 求和
	for i := range nums {
		res += nums[i]
	}
	return
}
