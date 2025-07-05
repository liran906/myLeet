// https://leetcode.cn/problems/permutations-ii/
// m

package backtrack

import "sort"

// 维护一个 dupe 数组记录同一层使用过的数字
// 也可以先排序再判断和前一个的值是否相同（空间更优），见方法 2
func permuteUnique_(nums []int) (ans [][]int) {
	var path []int

	// 标记每个元素是否在当前路径中被使用（基于索引）
	var used = make([]bool, len(nums))

	var bt func()
	bt = func() {
		if len(path) == len(nums) {
			// 得到一个完整的不重复排列
			ans = append(ans, append([]int{}, path...))
			return
		}

		// 本层搜索使用的数值（用于防止当前递归层内重复使用相同数字）
		// 数值范围在 [-10, 10] 映射为索引 [0, 20]
		dupe := make([]bool, 21)

		for i := 0; i < len(nums); i++ {
			// 如果这个索引的数字已经在当前 path 中使用过，跳过
			if used[i] {
				continue
			}

			// 如果这个数值在本层已经使用过了，跳过（防止同一层重复）
			if dupe[nums[i]+10] {
				continue
			}

			// 标记当前值在本层已用
			dupe[nums[i]+10] = true

			// 标记当前索引的值已用
			used[i] = true
			path = append(path, nums[i])

			// 进入下一层
			bt()

			// 回溯
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	bt()
	return
}

// 方法 2 先排序再判断和前一个的值是否相同（空间更优）
func permuteUnique(nums []int) (ans [][]int) {
	// 为了去重，我们先将数组排序。
	// 这样相同的数字就会挨在一起，便于剪枝处理。
	sort.Ints(nums)

	var path []int                  // 当前构造的排列路径
	used := make([]bool, len(nums)) // 标记每个元素是否被使用过（按索引）

	var bt func()
	bt = func() {
		// 如果当前排列路径已经包含了全部数字，说明构造完成
		if len(path) == len(nums) {
			// 将 path 的一个副本加入答案
			ans = append(ans, append([]int{}, path...))
			return
		}

		// 遍历每一个数字尝试加入 path 中
		for i := 0; i < len(nums); i++ {
			// 如果当前数字已经在 path 中被使用了，就跳过
			if used[i] {
				continue
			}

			// ====== 剪枝关键点 ======
			// 如果当前数字和前一个数字相同（nums[i] == nums[i-1]），
			// 且前一个数字在当前路径中没有被使用（!used[i-1]），
			// 说明我们是在“同一层”做选择，
			// 为了避免重复排列，跳过当前数字
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			// 做选择
			used[i] = true
			path = append(path, nums[i])

			// 进入下一层递归
			bt()

			// 撤销选择（回溯）
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	// 启动回溯
	bt()
	return
}
