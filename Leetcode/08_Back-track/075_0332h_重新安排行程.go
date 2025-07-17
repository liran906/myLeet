// https://leetcode.cn/problems/reconstruct-itinerary/description/
// h

package backtrack

import (
	"sort"
)

// 实现没问题，但
// 这个方法会超时
func findItinerary_oft(tickets [][]string) []string {
	// 构建邻接表，每个起点 maps 到若干终点和它们的使用次数
	// {"JFK" : {"MUC" : 1, "SFO" : 1}, ...}
	tmap := map[string]map[string]int{}
	for _, t := range tickets {
		from, to := t[0], t[1]
		if _, exists := tmap[from]; !exists {
			tmap[from] = map[string]int{}
		}
		tmap[from][to]++
	}

	// 回溯路径起点永远是 "JFK"
	var path = []string{"JFK"}

	// 保存所有合法路径（符合机票使用完、从 JFK 出发的要求）
	var allPaths [][]string

	// 定义回溯函数
	var backtrack func(string)
	backtrack = func(curr string) {
		// 如果 path 长度刚好是所有票数+1（即刚好使用所有票），记录路径
		if len(path) == len(tickets)+1 {
			// 注意要拷贝 path，否则会被后续修改
			allPaths = append(allPaths, append([]string{}, path...))
			return
		}

		// 遍历从当前城市出发的所有目的地（无序）
		for next, count := range tmap[curr] {
			if count > 0 {
				// 使用这张票
				tmap[curr][next]--
				path = append(path, next)

				backtrack(next)

				// 回溯：恢复状态
				path = path[:len(path)-1]
				tmap[curr][next]++
			}
		}
	}

	// 启动回溯
	backtrack("JFK")

	// 如果存在多个路径，选出字典序最小的那一个
	if len(allPaths) > 1 {
		sort.Slice(allPaths, func(i, j int) bool {
			a, b := allPaths[i], allPaths[j]
			for k := 0; k < len(a); k++ {
				if a[k] != b[k] {
					return a[k] < b[k]
				}
			}
			return false // 相同则不调换
		})
	}

	// 返回字典序最小的那条完整路径
	return allPaths[0]
}

// 标准回溯，但还是超时
func findItinerary_oft2(tickets [][]string) []string {
	// 1. 用 map 构建图：每个出发地对应一个目的地列表
	graph := map[string][]string{}
	for _, ticket := range tickets {
		graph[ticket[0]] = append(graph[ticket[0]], ticket[1])
	}

	// 2. 将每个出发地的目的地列表排序，确保回溯按字典序走
	for from := range graph {
		sort.Strings(graph[from])
	}

	// 3. 用 map[from] -> []bool 标记每一条边是否使用过
	used := map[string][]bool{}
	for from, dests := range graph {
		used[from] = make([]bool, len(dests))
	}

	var result []string
	var path []string

	var dfs func(string) bool
	dfs = func(from string) bool {
		path = append(path, from)

		// ✅ 如果用完所有票（走过 tickets 条边），返回成功
		if len(path) == len(tickets)+1 {
			result = append([]string{}, path...) // 复制一份
			return true
		}

		for i, to := range graph[from] {
			if !used[from][i] {
				used[from][i] = true
				if dfs(to) {
					return true // ✅ 找到有效路径直接返回，不再回溯
				}
				used[from][i] = false // ❌ 否则回退
			}
		}

		// 回溯
		path = path[:len(path)-1]
		return false
	}

	dfs("JFK")
	return result
}

// https://blog.golir.top/article/11
func findItinerary(tickets [][]string) []string {
	var (
		m   = map[string][]string{} // 邻接表：出发地 -> 目的地数组
		res []string                // 最终结果
	)

	// 构建邻接表
	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		m[src] = append(m[src], dst)
	}

	// 保证目的地按字典序排序（先访问字典序小的）
	for key := range m {
		sort.Strings(m[key])
	}

	var dfs func(curr string)
	dfs = func(curr string) {
		// 当前出发地有未使用的票
		for {
			if v, ok := m[curr]; !ok || len(v) == 0 {
				break
			}

			// 每次都取字典序最小的目的地（第一个）
			next := m[curr][0]

			// 从邻接表中删除这张票（边）
			m[curr] = m[curr][1:]

			dfs(next)
		}

		// 回溯时记录路径（post-order，逆序）
		res = append(res, curr)
	}

	dfs("JFK")

	// 结果是逆序的，需要翻转
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return res
}
