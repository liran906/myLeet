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

	// 逐步构建 dp 表，从小问题构建到大问题
	for i := 2; i <= n; i++ {
		// 枚举以 j 作为根节点的所有可能
		for j := 1; j <= i; j++ {
			// 左子树节点数为 j-1，右子树节点数为 i-j
			// 两者组合的乘积即为以 j 为根时的组合数
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	// 返回由 n 个节点组成的二叉搜索树的总数
	return dp[n]
}
