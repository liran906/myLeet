// https://leetcode.cn/problems/permutations/description/
// m

package backtrack

func permute(nums []int) (ans [][]int) {
	var path []int
	// 用于标记某个值是否被使用过
	// 根据题目说明，nums[i] ∈ [-10, 10]，所以共 21 个可能值
	// 使用值 +10 将其映射为索引 0 到 20
	var used = make([]bool, 21)

	var bt func()
	bt = func() {
		// 如果 path 的长度等于 nums 长度，说明生成了一个完整排列
		if len(path) == len(nums) {
			// 注意要拷贝 path，否则后续修改会影响已添加的结果
			ans = append(ans, append([]int{}, path...))
			return
		}

		// 遍历每一个数字
		for i := 0; i < len(nums); i++ {
			// 如果该值已经被使用（通过值而非索引判断），则跳过
			if used[nums[i]+10] {
				continue
			}

			// 标记当前值已使用
			used[nums[i]+10] = true
			path = append(path, nums[i])

			// 进入下一层递归
			bt()

			// 回溯：恢复状态
			used[nums[i]+10] = false
			path = path[:len(path)-1]
		}
	}

	// 从空排列开始搜索
	bt()
	return
}
