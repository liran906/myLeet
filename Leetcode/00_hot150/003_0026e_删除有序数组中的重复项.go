// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150

package hot150

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	var j int
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
