// https://leetcode.cn/problems/gas-station/description/
// m

package greedy

// 暴力 O(n^2)
// 遍历每一个加油站为起点的情况，模拟一圈。
// 如果跑了一圈，中途没有断油，而且最后油量大于等于0，说明这个起点是ok的。
func canCompleteCircuit_bf(gas []int, cost []int) int {
	n := len(gas)
	for i := range n {
		cgas := 0
		for j := 0; j < n && cgas >= 0; j++ {
			cur := (i + j) % n
			cgas += gas[cur] - cost[cur]
		}
		if cgas >= 0 {
			return i
		}
	}
	return -1
}

// 贪心解法：O(n)
// 如果总油量 < 总消耗，不可能跑完一圈。
// 否则，只需找到一个合适的起点即可。
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	curSum := 0   // 当前从 start 到 i 的净剩余油量
	totalSum := 0 // 所有站点的净剩余油量总和
	start := 0    // 起始加油站候选

	for i := range n {
		gain := gas[i] - cost[i]
		totalSum += gain
		curSum += gain

		// 当前起点无法支撑到下一个站点，重置起点
		if curSum < 0 {
			curSum = 0
			start = i + 1 // 重新选择下一个站点为起点
		}
	}

	// 只有当总剩余油量 >= 0，才能绕一圈成功
	if totalSum >= 0 {
		return start
	}
	return -1
}
