// https://leetcode.cn/problems/merge-intervals/description/
// m

package greedy

import "sort"

func merge(intervals [][]int) (merged [][]int) {
	// 按起始时间升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化当前合并区间的起点和终点
	currentStart := intervals[0][0]
	currentEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		nextStart := intervals[i][0]
		nextEnd := intervals[i][1]

		if nextStart > currentEnd {
			// 没有重叠，保存当前合并区间，并开始新的合并区间
			merged = append(merged, []int{currentStart, currentEnd})
			currentStart, currentEnd = nextStart, nextEnd
		} else {
			// 有重叠，更新合并区间的右边界
			currentEnd = max(currentEnd, nextEnd)
		}
	}

	// 添加最后一个合并区间
	merged = append(merged, []int{currentStart, currentEnd})
	return
}
