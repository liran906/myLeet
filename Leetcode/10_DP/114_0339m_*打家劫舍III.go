package dp

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 定义全局哈希表：缓存每个子树的最大收益，避免重复计算
var mymap = make(map[*TreeNode]int)

// rob 返回以 root 为根的二叉树能偷到的最大金额
func rob(root *TreeNode) int {
	// 空节点收益为 0
	if root == nil {
		return 0
	}

	// 如果这个节点的结果已经计算过，直接返回缓存值
	if val, ok := mymap[root]; ok {
		return val
	}

	// 叶子节点（没有左右孩子），最大收益就是偷它本身
	if root.Left == nil && root.Right == nil {
		mymap[root] = root.Val
		return mymap[root]
	}

	// 方案1：偷当前节点
	//  偷当前节点 → 左右孩子不能偷 → 只能偷左右孩子的左右孩子（即隔一层）
	res1 := root.Val
	if root.Left != nil {
		res1 += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		res1 += rob(root.Right.Left) + rob(root.Right.Right)
	}

	// 方案2：不偷当前节点
	//  不偷当前节点 → 可以偷左右孩子（直接下一层）
	res2 := rob(root.Left) + rob(root.Right)

	// 取两种方案的较大值，并缓存结果
	mymap[root] = max(res1, res2)
	return mymap[root]
}

// 更好的方法，树DP
func rob_dp(root *TreeNode) int {
	notRob, rob := dfs(root)
	return max(notRob, rob)
}

// 返回两个值：
// notRob → 当前节点不偷的最大收益
// rob    → 当前节点偷的最大收益
func dfs(node *TreeNode) (notRob int, rob int) {
	if node == nil {
		return 0, 0
	}

	leftNotRob, leftRob := dfs(node.Left)
	rightNotRob, rightRob := dfs(node.Right)

	// 不偷当前节点：左右孩子可偷可不偷，取最大值
	notRob = max(leftNotRob, leftRob) + max(rightNotRob, rightRob)

	// 偷当前节点：左右孩子都不能偷
	rob = node.Val + leftNotRob + rightNotRob

	return notRob, rob
}
