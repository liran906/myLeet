// https://kamacoder.com/problempage.php?pid=1052

package dp

import "fmt"

func wholeBackpack() {
	var n, v int
	fmt.Scan(&n, &v)

	imap := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, v)
		for j := 0; j < 2; j++ {
			fmt.Scan(&row[j])
		}
		imap[i] = row
	}

	var dp = make([]int, v+1)
	for i := 0; i < len(imap); i++ {
		weight, value := imap[i][0], imap[i][1]
		// 完全背包比 01 背包关键区别就是 j 正序遍历（压缩的话）
		for j := weight; j <= v; j++ {
			dp[j] = max(dp[j], dp[j-weight]+value)
		}
	}
	fmt.Println(dp[v])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
