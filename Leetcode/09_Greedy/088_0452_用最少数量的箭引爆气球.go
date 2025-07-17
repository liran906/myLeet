// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/description/
// m

package greedy

import "sort"

func findMinArrowShots(points [][]int) (count int) {
	// 将所有气球按照“左边界”从小到大排序（也可以按右边界排序，逻辑不同但思路相似）
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	// 当前“瞄准点”设置为第一个气球的右边界（也就是该箭最远能打到哪里）
	cur := points[0][1]
	count = 1 // 题目说最少一个气球，那么最少也需要 1 箭
	for i := 1; i < len(points); i++ {
		left, right := points[i][0], points[i][1]
		if cur < left {
			// 当前气球和上一组气球没有交集（新气球左边界都大于瞄准点）
			// 需要重新射一箭
			cur = right
			count++
		} else {
			// 当前气球和上一组气球有交集，可以被同一箭射中，省下一箭
			// 更新瞄准点为两个气球右边界的较小值（保证后续的交集尽可能延续）
			cur = min(cur, right)
		}
	}
	return
}
