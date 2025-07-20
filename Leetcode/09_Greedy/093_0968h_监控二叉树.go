// https://leetcode.cn/problems/binary-tree-cameras/description/
// h

package greedy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) (count int) {
	// 定义递归函数，后序遍历二叉树
	var helper func(*TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 后续遍历，因为要按照 叶节点 -> 根节点的顺序去贪心
		// 所以先递归左右子树，再写核心逻辑（根）
		helper(node.Left)
		helper(node.Right)

		// 核心逻辑：
		// 状态定义：
		// 0：当前节点没有被覆盖
		// 1：当前节点被覆盖但没有摄像头
		// 2：当前节点安装了摄像头

		// 如果任一子节点没有被覆盖，则当前节点必须安装摄像头
		if (node.Left != nil && node.Left.Val == 0) || (node.Right != nil && node.Right.Val == 0) {
			count++
			node.Val = 2 // 当前节点安装摄像头

			// 被摄像头覆盖
			if node.Left != nil {
				node.Left.Val = 1
			}
			if node.Right != nil {
				node.Right.Val = 1
			}

		} else if (node.Left != nil && node.Left.Val == 2) || (node.Right != nil && node.Right.Val == 2) {
			// 如果任一子节点有摄像头，则当前节点被覆盖
			node.Val = 1
		}
		// 否则 node.Val 默认为 0，等待上层处理
	}

	// 开始递归
	helper(root)

	// 根节点可能没有被覆盖，单独处理
	if root.Val == 0 {
		count++
	}

	return
}
