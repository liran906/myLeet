// https://leetcode.cn/problems/monotone-increasing-digits/description/
// m

package greedy

import "strconv"

func monotoneIncreasingDigits(n int) int {
	nstr := []byte(strconv.Itoa(n))
	mark := len(nstr) // 标记从哪里开始填 '9'

	// 从后往前遍历，发现下降点，就标记并将前一位减1
	for i := len(nstr) - 1; i > 0; i-- {
		if nstr[i-1] > nstr[i] {
			mark = i
			nstr[i-1]-- // 比如 352 前一位减 1，当前位变 9，最后是 349
		}
	}

	// 将标记之后的位置全设为 '9'
	for i := mark; i < len(nstr); i++ {
		nstr[i] = '9'
	}

	// 转回整数
	ans, _ := strconv.Atoi(string(nstr))
	return ans
}
