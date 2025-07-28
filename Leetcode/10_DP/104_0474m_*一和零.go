// https://leetcode.cn/problems/ones-and-zeroes/description/
// m

package dp

// 这个方法错在把 n 和 m 分别考虑了，但他们需要共同约束才行
func findMaxForm_wrong(strs []string, m int, n int) int {
	var mCount, nCount []int
	for _, str := range strs {
		var mcount, ncount int
		for j := range str {
			if str[j] == '0' {
				mcount++
			} else {
				ncount++
			}
		}
		mCount = append(mCount, mcount)
		nCount = append(nCount, ncount)
	}

	// 把子集的长度看作 value, weight 分别是 mCount 和 nCount
	// 要实现 weight 不超标前提下，最大化 value
	dpm := make([]int, m+1)
	dpn := make([]int, n+1)
	for i := range strs {
		for jm := m; jm >= 0; jm-- {
			if mCount[i] <= jm {
				dpm[jm] = max(dpm[jm], dpm[jm-mCount[i]]+1)
			}
		}
		for jn := n; jn >= 0; jn-- {
			if nCount[i] <= jn {
				dpn[jn] = max(dpn[jn], dpn[jn-nCount[i]]+1)
			}
		}
	}
	return min(dpm[m], dpn[n])
}

/*
问题建模为：二维容量限制的 0-1 背包问题（Dual-Dimensional 0/1 Knapsack）：

状态定义（原始思路）：
- 定义三维数组 dp[k][i][j] 表示：
  使用前 k 个字符串，消耗 i 个 0、j 个 1 所能拼出的最大子集长度。

状态转移方程：
- 不选第 k 个字符串：dp[k][i][j] = dp[k-1][i][j]
- 选第 k 个字符串（0 数量为 zeros，1 数量为 ones）：
  若 i ≥ zeros 且 j ≥ ones：
  dp[k][i][j] = max(dp[k-1][i][j], dp[k-1][i-zeros][j-ones] + 1)

**空间压缩**（和 2d-dp 压缩到 1d-dp 相同的道理）：
- 由于第 k 层只依赖第 k-1 层，可以将三维数组压缩为二维：
  dp[i][j] 表示在当前物品集合下最多能拼出多少个字符串
- 注意要从后往前遍历 i 和 j，以避免同一轮中重复使用某个字符串（0-1 背包）

最终目标：
- dp[m][n] 即为在不超过 m 个 0 和 n 个 1 的限制下，最多可以选择多少个字符串

解题思路：
本题是一个「二维容量限制」的 0-1 背包问题：
- 每个字符串只能使用一次（0-1 背包）
- 每个字符串消耗一定数量的 0 和 1（二维体积：zeros 和 ones）
- 每个字符串的价值都是 1（我们要求的是最多能选多少个字符串）

具体做法如下：
1. 初始化一个二维 DP 数组 dp[i][j] 表示：
   在不超过 i 个 '0' 和 j 个 '1' 的限制下，最多可以选多少个字符串。
2. 遍历每个字符串，统计它的 0 和 1 数量（即「重量」）
3. 从后往前更新 dp 表，保证每个字符串只被用一次（0-1 背包的特性）
4. 状态转移：
   dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones] + 1)

最终答案就是 dp[m][n]，即在最大允许 m 个 0 和 n 个 1 下，最多可以选出多少个字符串。
*/

// findMaxForm 返回能从 strs 中选出最多的子集，使得其中包含的 0 的个数不超过 m，1 的个数不超过 n
func findMaxForm(strs []string, m int, n int) int {
	// 初始化一个二维 DP 表，dp[i][j] 表示使用 i 个 0 和 j 个 1 所能组成的最大子集长度
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 遍历每一个字符串，视作一个「物品」
	for _, str := range strs {
		// 统计当前字符串中 0 和 1 的数量 —— 也就是这个物品的两个「重量」
		zeros, ones := 0, 0
		for _, ch := range str {
			if ch == '0' {
				zeros++
			} else {
				ones++
			}
		}

		// 从大到小遍历 dp 表：逆序保证每个字符串只能被使用一次（0-1 背包）
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				// 状态转移方程：
				// 不选当前字符串：dp[i][j]
				// 选当前字符串：dp[i-zeros][j-ones] + 1
				dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones]+1)
			}
		}
	}

	// 最终结果是使用 m 个 0 和 n 个 1 所能组成的最大子集数
	return dp[m][n]
}
