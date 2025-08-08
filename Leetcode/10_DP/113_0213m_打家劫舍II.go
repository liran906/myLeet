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
