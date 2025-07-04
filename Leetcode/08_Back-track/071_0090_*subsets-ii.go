// https://leetcode.cn/problems/subsets-ii/
// m

package backtrack

import "sort"

func subsetsWithDup(nums []int) (ans [][]int) {
	// 先排序，使得重复元素相邻，便于后续判断是否重复
	sort.Ints(nums)

	var path []int
	var bt func(index int)
	bt = func(index int) {
		// 子集问题与组合问题不同：每个节点都是一个子集（不是只在叶子节点收集）
		// 对比 40 题
		ans = append(ans, append([]int{}, path...))

		// 遍历当前位置之后的每一个数（从 index 开始）
		for i := index; i < len(nums); i++ {

			// 去重逻辑（关键）：
			// 如果当前数 nums[i] 与前一个 nums[i-1] 相等，且 i > index，
			// 则说明这是同一层（for 循环中）出现的重复数 —— 跳过
			if i > index && nums[i] == nums[i-1] {
				continue // 跳过重复，保证同一层不会选择同一个数两次
			}

			// 对比 40 题
			// 不同层出现重复数是允许的！
			// 比如：[1,2] 和 [1,2,2] —— 第二个 2 是在下一层递归中选的
			// 因为递归是逐层向下推进的，每一层都会有自己的 for 循环
			// 即使是相同的数，只要不是在同一 for 层中出现，就不视为重复

			path = append(path, nums[i])
			bt(i + 1)
			path = path[:len(path)-1]
		}
	}
	bt(0)
	return ans
}
