// https://leetcode.cn/problems/candy/description/
// h

package greedy

// O(n)左右两次遍历
func candy(ratings []int) (res int) {
	// nums[i] 表示第 i 个孩子在基础 1 颗糖外的额外糖果数量（初始为 0）
	var nums = make([]int, len(ratings))

	// 第一次遍历：从左到右
	// 如果右边孩子评分高于左边，则 nums[i] = nums[i-1] + 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			nums[i] = nums[i-1] + 1
		}
	}

	// 第二次遍历：从右到左，同时累加糖果总数
	for i := len(ratings) - 2; i >= 0; i-- {
		// 如果左边孩子评分高于右边，则 nums[i] 至少为 nums[i+1] + 1
		if ratings[i] > ratings[i+1] {
			nums[i] = max(nums[i], nums[i+1]+1)
		}

		// 当前孩子的糖果 = nums[i] + 1（因为每人至少一颗）
		res += nums[i] + 1
	}

	// 最后加上最后一个孩子的糖果（由于上面的循环没有加到）
	return res + nums[len(nums)-1] + 1
}
