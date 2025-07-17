// https://leetcode.cn/problems/non-overlapping-intervals/description/
// m

package greedy

import "sort"

func eraseOverlapIntervals(intervals [][]int) (count int) {
	// 按照每个区间的起始点升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化“前一个区间”的右边界
	rightMost := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		left, right := intervals[i][0], intervals[i][1]

		if rightMost > left {
			// 当前区间与前一个区间重叠
			count++ // 删除一个重叠区间（只记录删除次数）

			// 贪心地保留结束位置更小的那个（重叠更少，有助于后续保留更多）
			rightMost = min(rightMost, right)
		} else {
			// 无重叠，更新右边界为当前区间的结束位置
			rightMost = right
		}
	}

	return
}
