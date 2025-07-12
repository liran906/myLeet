// https://leetcode.cn/problems/wiggle-subsequence/
// m

package greedy

// 贪心 O(n)
func wiggleMaxLength(nums []int) int {
	// 长度 0 或 1 的数组，直接返回长度
	if len(nums) <= 1 {
		return len(nums)
	}

	// 跳过一开始连续相等的元素
	index := 1
	for index < len(nums) && nums[index] == nums[index-1] {
		index++
	}

	// 如果所有元素都相等，只有一个值能构成 wiggle 序列
	if index == len(nums) {
		return 1
	}

	// 至少有两个不相等元素，初始长度为 2
	count := 2
	// 初始差值（用于判断正负跳跃）
	dif := nums[index] - nums[index-1]

	// 从第 index+1 个数开始遍历
	for i := index + 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue // 相等的数不能构成 wiggle
		}

		// 如果当前差值方向与前一个差值方向相反
		if dif*(nums[i]-nums[i-1]) < 0 {
			count++   // 成功“拐弯”一次，计数+1
			dif *= -1 // 翻转差值方向（正变负，负变正）
		}
	}
	return count
}

// 优化写法
func wiggleMaxLength_(nums []int) int {
	// 若数组长度小于 2，则最长摆动子序列长度就是其本身长度
	if len(nums) < 2 {
		return len(nums)
	}

	// 计算前两个数之间的初始差值
	prevDiff := nums[1] - nums[0]

	// 若初始差值为 0，则初始长度为 1；否则为 2
	count := 1
	if prevDiff != 0 {
		count = 2
	}

	// 从第三个数开始遍历
	for i := 2; i < len(nums); i++ {
		// 当前差值
		diff := nums[i] - nums[i-1]

		// 满足摆动条件：当前差值和前一个差值方向相反
		if (diff > 0 && prevDiff <= 0) || (diff < 0 && prevDiff >= 0) {
			count++         // 成功“拐弯”，计数加一
			prevDiff = diff // 更新前一个差值为当前差值
		}
		// 如果 diff == 0 或方向未改变，则跳过当前元素
	}

	return count
}
