// https://leetcode.cn/problems/unique-binary-search-trees/description/
// m

package dp

// 这题的解析看代码随想录
// https://programmercarl.com/0096.%E4%B8%8D%E5%90%8C%E7%9A%84%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91.html
func numTrees(n int) int {
	// dp[i] 表示由 i 个节点组成的二叉搜索树的种类数
	dp := make([]int, n+1)

	// 空树和单节点树都只有一种形式
	dp[0], dp[1] = 1, 1

	// 构建 dp 表：dp[i] 表示由 i 个节点组成的 BST 的种类数
	for i := 2; i <= n; i++ {
		// 枚举所有可能的根节点 j（1 到 i）
		// 对于每个 j：
		//   左子树有 j-1 个节点，右子树有 i-j 个节点
		//   左右子树的组合方式为 dp[j-1] * dp[i-j]
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	// 返回由 n 个节点组成的二叉搜索树的总数
	return dp[n]
}
