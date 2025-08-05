// https://kamacoder.com/problempage.php

package dp

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	dp := make([]int, n+1)
	dp[0] = 1

	for j := 0; j <= n; j++ {
		for i := 1; i <= m; i++ {
			if j >= i {
				dp[j] = dp[j] + dp[j-i]
			}
		}
	}
	fmt.Println(dp[n])
}
