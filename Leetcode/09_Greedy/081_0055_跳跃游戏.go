//https://leetcode.cn/problems/jump-game/
//m

package greedy

// 暴力递归
func canJump_bt(nums []int) bool {
	var helper func(int) bool
	helper = func(cur int) bool {
		if cur == len(nums)-1 {
			return true // 成功跳到最后一格
		}
		if nums[cur] == 0 {
			return false // 当前无法前进
		}

		// 尝试所有能跳的步数
		for i := 1; i <= nums[cur] && cur+i < len(nums); i++ {
			if helper(cur + i) {
				return true
			}
		}
		return false
	}
	return helper(0)
}

// 记忆 O n^2
func canJump_btm(nums []int) bool {
	// 用 table 记录是否可以到达
	var table = map[int]bool{}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] >= len(nums)-1-i {
			table[i] = true
			continue
		}
		for j := nums[i]; j > 0; j-- {
			if table[i+j] == true {
				table[i] = true
				break
			}
		}
	}
	return table[0]
}

// 贪心 On 从右到左
func canJump(nums []int) bool {
	// 初始化最远可以跳到终点的“最左边”位置
	leftBoundary := len(nums) - 1

	// 从倒数第二个元素开始向前遍历
	for i := len(nums) - 2; i >= 0; i-- {
		// 如果当前位置可以跳到或超过当前的 leftBoundary
		if i+nums[i] >= leftBoundary {
			// 更新 leftBoundary 为当前位置
			// 意味着从这个位置也能跳到终点
			leftBoundary = i
		}
	}
	// 如果最终 leftBoundary 被推进到了 0
	// 表示从起点就可以跳到终点
	return leftBoundary == 0
}

// 贪心 On 从左到右
func canJump_2(nums []int) bool {
	// rightBoundary 表示从起点出发目前能跳到的最远位置
	rightBoundary := nums[0]

	// 从第 1 个位置开始判断是否可达
	for i := 1; i < len(nums); i++ {
		// 如果当前位置 i 已经超出了能到达的最远位置，返回 false
		if i > rightBoundary {
			return false
		}

		// 更新能跳到的最远位置
		rightBoundary = max(rightBoundary, i+nums[i])
	}

	// 如果整个数组都可以遍历完，说明可以到达最后一个位置
	return true
}
