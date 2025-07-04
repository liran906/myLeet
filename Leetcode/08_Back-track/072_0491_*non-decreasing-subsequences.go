// https://leetcode.cn/problems/non-decreasing-subsequences/
// m

package backtrack

import "math"

// 这题有同层的逻辑（不重复），同层逻辑在 for 外面增加记录，并且不传递进入递归
// 也有不同层的逻辑（非递减），非同层逻辑要靠递归参数传递到后续层

func findSubsequences(nums []int) (ans [][]int) {
	var path []int

	// 回溯函数：
	// index 表示当前从哪个位置开始选数
	// last 表示当前路径中最后一个数，用于判断是否递增
	var bt func(index int, last int)
	bt = func(index, last int) {
		// ✅ 当 path 长度 ≥ 2，说明形成了一个合法子序列，可以加入结果中
		if len(path) > 1 {
			ans = append(ans, append([]int{}, path...))
		}

		// ✅ 同层去重：记录当前递归层级中使用过的元素值，避免重复使用
		seen := map[int]struct{}{}

		for i := index; i < len(nums); i++ {
			// ✳️ 跳过不递增的情况，保证子序列递增
			if nums[i] < last {
				continue
			}

			// ❗ 如果当前层已经使用过这个数字，跳过，防止同一层重复
			if _, exists := seen[nums[i]]; exists {
				continue
			}

			// ✅ 标记当前值为“本层已用”，避免后续同层重复使用
			seen[nums[i]] = struct{}{}

			// 回溯
			path = append(path, nums[i])
			bt(i+1, nums[i])
			path = path[:len(path)-1]
		}
	}

	// 初始调用，使用 math.MinInt 允许第一个数无条件加入
	bt(0, math.MinInt)

	return
}

// 小优化：用数组代替 map
func findSubsequences_2(nums []int) (ans [][]int) {
	var path []int

	// 回溯函数：index 是当前搜索起点，last 是上一个选择的值
	var bt func(int, int)
	bt = func(index, last int) {
		// 当 path 中元素个数 ≥ 2 时，即为合法递增子序列，加入结果集
		if len(path) > 1 {
			ans = append(ans, append([]int{}, path...))
		}

		// 创建数组代替哈希集合，记录本层使用过的数字，避免重复分支
		// 题目已说明 nums[i] ∈ [-100, 100]，所以 +100 映射到 [0,200]
		// 虽然都是常数级复杂度，但比用 map 更快更高效（没有 hash 计算、重复、扩容）
		seen := make([]int, 201)

		for i := index; i < len(nums); i++ {
			// 不是递增（非严格递增也可以，允许相等）
			if nums[i] < last {
				continue
			}

			// 当前层已经使用过 nums[i]，避免重复选择
			if seen[nums[i]+100] == 1 {
				continue
			}

			// 标记当前层已使用
			seen[nums[i]+100] = 1

			// 回溯
			path = append(path, nums[i])
			bt(i+1, nums[i])
			path = path[:len(path)-1]
		}
	}

	// 初始化 last 为最小整数，以便首次 nums[i] < last 不会发生
	bt(0, math.MinInt)
	return
}
