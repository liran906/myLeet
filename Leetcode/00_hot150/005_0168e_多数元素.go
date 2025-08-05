// https://leetcode.cn/problems/majority-element/description/?envType=study-plan-v2&envId=top-interview-150
package hot150

// 这题遇到了就正常用排序或者hash记录数量即可。

// 算法思想说明（Boyer–Moore 投票算法） O(n) O(1)
//  1. 基本假设： 多数元素在数组中出现次数 > ⌊n/2⌋，所以它的“票数”最终一定 > 0。
//  2. 每次遇到当前候选者的票，票数加一；否则减一。
//  3. 票数归零时，重新换一个候选者。
//  4. 因为多数元素的个数比其他所有元素之和还多，它一定会赢得最后的“投票”。
func majorityElement(nums []int) int {
	votes := 0 // 记录当前候选元素的“票数”
	cur := 0   // 当前候选的多数元素

	// 遍历整个数组
	for i := 0; i < len(nums); i++ {
		// 如果当前没有候选者（票数归零），就设定当前元素为候选者
		if votes == 0 {
			cur = nums[i] // 切换候选人
		}

		// 如果当前元素等于候选者，则票数+1；否则票数-1
		if nums[i] == cur {
			votes++ // 支持票
		} else {
			votes-- // 反对票
		}
	}

	// 最终返回的 cur 就是“多数元素”
	return cur
}
