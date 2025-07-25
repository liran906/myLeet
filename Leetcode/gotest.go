package main

import (
	"fmt"
	"sort"
)

func main() {
	b := canPartition([]int{1, 5, 5, 11})
	fmt.Println(b)
}

func canPartition(nums []int) bool {
	sort.Ints(nums)
	fmt.Println(nums)
	l, r := 0, len(nums)-1
	sum := 0
	for l < r {
		fmt.Printf("l=%d r=%d sum=%d\n", l, r, sum)
		if sum < 0 {
			sum += nums[l]
			l++
		} else {
			sum -= nums[r]
			r--
		}
	}
	fmt.Println(sum)
	return sum == 0
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func testMT() {
	t1 := &TreeNode{Val: 5}
	t2 := &TreeNode{Val: 4, Right: t1}
	t3 := &TreeNode{Val: 3}
	t4 := &TreeNode{Val: 2, Left: t2, Right: t3}
	t5 := &TreeNode{Val: 1}
	t6 := &TreeNode{Val: 0, Left: t4, Right: t5} // Tree 1

	t7 := &TreeNode{Val: 9}
	t8 := &TreeNode{Val: 8, Left: t7}
	t9 := &TreeNode{Val: 7}
	t10 := &TreeNode{Val: 6, Right: t9}
	t11 := &TreeNode{Val: 10, Left: t8, Right: t10} // Tree 2

	merged := mergeTrees(t6, t11)
	fmt.Println(merged)
}

type TreeNodes struct {
	tn  *TreeNode
	tn1 *TreeNode
	tn2 *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root := &TreeNode{Val: root1.Val + root2.Val}
	queue := []TreeNodes{
		{tn: root, tn1: root1, tn2: root2},
	}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		node, node1, node2 := top.tn, top.tn1, top.tn2

		node.Left = nodeSum(node1.Left, node2.Left)
		if node.Left != nil {
			queue = append(queue, TreeNodes{
				tn: node.Left, tn1: node1.Left, tn2: node2.Left,
			})
		}

		node.Right = nodeSum(node1.Right, node2.Right)
		if node.Right != nil {
			queue = append(queue, TreeNodes{
				tn: node.Right, tn1: node1.Right, tn2: node2.Right,
			})
		}
	}
	return root
}

func nodeSum(node1, node2 *TreeNode) *TreeNode {
	if node1 != nil && node2 != nil {
		return &TreeNode{Val: node1.Val + node2.Val}
	}
	if node1 != nil {
		return &TreeNode{Val: node1.Val}
	}
	if node2 != nil {
		return &TreeNode{Val: node2.Val}
	}
	return nil
}

func testFI() {
	tickets := [][]string{{"JFK", "a"}, {"JFK", "b"}, {"c", "JFK"}, {"b", "c"}}
	res := findItinerary(tickets)
	fmt.Println(res)
}

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
			fmt.Println("current: ", curr)

			if v, ok := m[curr]; !ok || len(v) == 0 {
				fmt.Println("reached end...")
				break
			}

			fmt.Println("next choices: ", m[curr])

			// 每次都取字典序最小的目的地（第一个）
			next := m[curr][0]

			// 从邻接表中删除这张票（边）
			m[curr] = m[curr][1:]

			dfs(next)
		}

		// 回溯时记录路径（post-order，逆序）
		fmt.Println("appending", curr)
		res = append(res, curr)
		fmt.Println("res after append: ", res)
	}

	dfs("JFK")

	// 结果是逆序的，需要翻转
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return res
}
