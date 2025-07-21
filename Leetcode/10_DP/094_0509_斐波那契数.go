package dp

var m = make(map[int]int)

// 递归+存储
func fib(n int) int {
	if n <= 1 {
		return n
	}
	if v, ok := m[n]; ok {
		return v
	}
	ans := fib(n-1) + fib(n-2)
	m[n] = ans
	return ans
}

// 正常 dp 写法
func fib_(n int) int {
	if n <= 1 {
		return n
	}
	dp := make(map[int]int)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
