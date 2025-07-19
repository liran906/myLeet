// https://leetcode.cn/problems/partition-labels/description/
// m

package greedy

func partitionLabels(s string) (ans []int) {
	// 记录每个字符最后出现的位置（下标）
	m := make([]int, 26) // 只考虑小写字母
	for i := 0; i < len(s); i++ {
		b := s[i]
		m[b-'a'] = i // 更新字符的最远位置
	}

	// 当前分区的结束下标，上一个分区的结束下标
	end, prev := -1, -1 // 因为逻辑是从上一个区间最后一个字母开始算，所以这里都要取 -1
	for i := 0; i < len(s); i++ {
		b := s[i]
		// 每个字符都可能延长当前分区的结束位置
		end = max(end, m[b-'a'])

		// 当遍历到当前分区的最后一个位置时，记录分区长度
		if i == end {
			ans = append(ans, end-prev)
			prev = end // 更新上一个分区的结尾
		}
	}
	return
}
