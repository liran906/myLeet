// https://leetcode.cn/problems/rotate-array/description/?envType=study-plan-v2&envId=top-interview-150
package hot150

// 原地反转的方法
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n // 避免 k > n 的情况

	// 1. 反转整个数组
	reverse(nums, 0, n-1)
	// 2. 反转前 k 个元素
	reverse(nums, 0, k-1)
	// 3. 反转剩余元素
	reverse(nums, k, n-1)
}

// 反转 nums[l:r] 区间
func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
