package dp

// bfs 超时
// uniquePaths 是计数问题，不是搜索最短路或是否能到达的问题，所以 不能剪枝
// 算法复杂度是指数级的
func uniquePaths_oft(m int, n int) (count int) {
	moves := [][]int{{1, 0}, {0, 1}}
	queue := [][]int{{0, 0}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		row, col := cur[0], cur[1]
		if row >= m || col >= n {
			continue
		}
		if row == m-1 && col == n-1 {
			count++
			continue
		}

		for _, m := range moves {
			next := []int{row + m[0], col + m[1]}
			queue = append(queue, next)
		}
	}
	return
}

// dp 写法1 不够优雅
func uniquePaths_(m int, n int) int {
	// 由于 slice 不能作为 map 的 key，所以这里把地图想象为一维数组
	// 按照 row -> col 顺序遍历，比如 m=2 c=4 的时候，就是一个长度为 8 的一维数组
	// 下标 1 就是（r=0,c=1）,下标 6 就是（r=1,c=2）
	dp := make([]int, m*n)
	dp[0] = 1
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if r == 0 && c == 0 {
				continue
			}
			index := r*n + c
			up := index - n
			left := index - 1
			if up < 0 {
				dp[index] = dp[left]
			} else if index%n == 0 {
				dp[index] = dp[up]
			} else {
				dp[index] = dp[up] + dp[left]
			}
		}
	}
	return dp[m*n-1]
}

// dp 写法 2 更好的写法！
func uniquePaths(m int, n int) int {
	// 使用一维数组表示整个网格，共 m 行 n 列，总长度 m*n
	dp := make([]int, m*n)
	dp[0] = 1 // 起点 (0,0) 的路径数为 1

	for r := range m {
		for c := range n {
			index := r*n + c  // 当前点的一维下标
			up := index - n   // 上方的格子
			left := index - 1 // 左侧的格子

			if r > 0 {
				dp[index] += dp[up] // 来自上方的路径数
			}
			if c > 0 {
				dp[index] += dp[left] // 来自左方的路径数
			}
		}
	}
	return dp[m*n-1] // 返回右下角的路径数
}
