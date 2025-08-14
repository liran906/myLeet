package hot150

import "sort"

func hIndex(citations []int) int {
	// 1. 升序排序，方便后面用 "有多少篇论文引用数 >= 当前论文引用数"
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] < citations[j]
	})

	res := 0
	// 2. 遍历每一篇论文（i 从 0 到 n-1）
	for i := range citations {
		// 当前论文的引用数
		// citations[i] 是升序排列的，所以 citations[i] 越靠后越大
		// cur 表示“至少有 cur 篇论文的引用数 >= citations[i]”
		cur := len(citations) - i

		// H 指数定义：至少有 h 篇论文的引用数 >= h
		// 所以候选值是 min(cur, citations[i])
		// 取所有候选值的最大值
		res = max(res, min(cur, citations[i]))
	}
	return res
}
