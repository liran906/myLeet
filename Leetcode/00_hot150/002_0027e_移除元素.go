// https://leetcode.cn/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150

package hot150

func removeElement(nums []int, val int) int {
	i, j := 0, 0
	count := 0
	for i < len(nums) {
		if nums[i] != val {
			count++
			nums[j] = nums[i]
			j++
		}
		i++
	}
	return count
}
