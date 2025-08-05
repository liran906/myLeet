// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/description/?envType=study-plan-v2&envId=top-interview-150
package hot150

func removeDuplicates_ii(nums []int) int {
	// 如果数组长度小于等于 2，说明不可能出现超过两个重复的情况，直接返回长度即可
	if len(nums) <= 2 {
		return len(nums)
	}

	// j 是慢指针，表示当前要放置元素的位置（第一个和第二个元素总是保留）
	j := 1

	// i 是快指针，从第三个元素开始扫描
	for i := 2; i < len(nums); i++ {
		// 如果当前元素 nums[i] ≠ nums[j-1]，说明它不是重复超过两次的值
		if nums[i] != nums[j-1] {
			j++               // 前移慢指针
			nums[j] = nums[i] // 将当前元素放入正确的位置
		}
		// 否则说明 nums[i] == nums[j-1]，而 nums[j-1] 一定等于 nums[j]（因为是排好序的）
		// 所以 nums[i] 是第 3 次或更多次重复，跳过它
	}

	// 最终返回去重后数组的长度（j 是索引，从 0 开始，长度应为 j + 1）
	return j + 1
}
