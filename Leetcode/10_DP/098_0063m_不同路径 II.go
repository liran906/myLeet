// https://leetcode.cn/problems/unique-paths-ii/description/
// m

package dp

// 一维数组的 dp
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)    // 行数
	n := len(obstacleGrid[0]) // 列数

	// 使用一维数组表示 m 行 n 列的 DP 表（压缩空间）
	dp := make([]int, m*n)

	dp[0] = 1 // 起点设为 1，表示从起点出发有一条路径

	for r := range m { // 遍历每一行
		for c := range n { // 遍历每一列
			index := r*n + c // 将二维下标 (r,c) 映射为一维下标 index

			// 如果当前位置是障碍物，路径数为 0（不可达）
			if obstacleGrid[r][c] == 1 {
				dp[index] = 0
				continue
			}

			// 如果不是第一行，则可以从上方走到当前位置
			if r > 0 {
				dp[index] += dp[index-n]
			}

			// 如果不是第一列，则可以从左方走到当前位置
			if c > 0 {
				dp[index] += dp[index-1]
			}
		}
	}

	// 返回终点位置的路径总数
	return dp[m*n-1]
}

// 二维数组形式，可读性更好一点
func uniquePathsWithObstacles_2d(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1
	for r := range m {
		for c := range n {
			if obstacleGrid[r][c] == 1 {
				dp[r][c] = 0
				continue
			}
			if r > 0 {
				dp[r][c] += dp[r-1][c]
			}
			if c > 0 {
				dp[r][c] += dp[r][c-1]
			}
		}
	}
	return dp[m-1][n-1]
}
