// https://leetcode.cn/problems/assign-cookies/description/
// e

package greedy

import "sort"

func findContentChildren(g []int, s []int) (count int) {
	// 如果没有饼干，直接返回 0
	if len(s) == 0 {
		return
	}

	// 将孩子的胃口和饼干大小都从小到大排序
	sort.Ints(g)
	sort.Ints(s)

	// 从最后一个（最大的）饼干开始尝试分配
	cookie := len(s) - 1

	// 从胃口最大的小孩开始尝试满足
	for child := len(g) - 1; child >= 0; child-- {
		// 如果当前饼干可以满足这个孩子
		if cookie >= 0 && s[cookie] >= g[child] {
			count++  // 满足一个孩子
			cookie-- // 用掉一块饼干
		}
		// 否则：当前饼干太小了，尝试满足下一个孩子（继续往前找更小胃口的）
	}

	return
}
