// https://leetcode.cn/problems/queue-reconstruction-by-height/description/
// m

package greedy

import (
	"container/list"
	"sort"
)

func reconstructQueue(people [][]int) (res [][]int) {
	// ✅ Step 1: 排序规则
	//   - 优先按身高 h 降序排列（因为高个子不会被矮个子影响）
	//   - 如果身高相同，则按 k 升序排列（k 小的排在前面）
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			// 身高相同，谁前面应有的人更少谁排前面
			return people[i][1] < people[j][1]
		}
		// 身高高的排前面
		return people[i][0] > people[j][0]
	})

	// ✅ Step 2: 插入每个人到目标队列中
	//   - 插入位置为他们的 k 值
	for _, p := range people {
		idx := p[1] // 插入位置为 k（前面要有 k 个大于等于 h 的人）

		// ✅ 插入操作：
		// 先扩容（末尾加个空位），再把插入位置后的元素右移一位
		// 最后将 p 放到 idx 位置上
		res = append(res, nil)       // 扩容
		copy(res[idx+1:], res[idx:]) // 元素右移
		res[idx] = p                 // 插入 p
	}

	return res
}

// 链表实现
func reconstructQueue_ll(people [][]int) [][]int {
	// 第一步：排序
	// 规则：
	// - 身高从高到低排（降序）
	// - 身高相同时，k 从小到大排（升序）
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	// 第二步：使用链表按顺序插入每个人
	// 这样插入不会打乱之前已插入的高个子位置
	l := list.New() // 创建一个双向链表

	for _, person := range people {
		position := person[1]      // 要插入的位置（即前方应该有几个人）
		node := l.PushBack(person) // 临时加到末尾，为了移动用

		// 从头开始遍历，找到目标位置
		cur := l.Front()
		for i := 0; i < position; i++ {
			cur = cur.Next()
		}

		// 移动到正确位置之前（cur 可能是 nil，即放到末尾）
		l.MoveBefore(node, cur)
	}

	// 第三步：把链表转回 [][]int 切片
	var res [][]int
	for e := l.Front(); e != nil; e = e.Next() {
		res = append(res, e.Value.([]int))
	}
	return res
}
