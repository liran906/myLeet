// https://leetcode.cn/problems/word-break/

package dp

// https://programmercarl.com/0139.%E5%8D%95%E8%AF%8D%E6%8B%86%E5%88%86.html
// 看下代码随想录的解析，这题要是排列，而不是组合，所以j在外层
func wordBreak(s string, wordDict []string) bool {
	// dp[j] 表示：s[0:j]（前 j 个字符）是否可以被 wordDict 中的单词拼出
	dp := make([]bool, len(s)+1)
	dp[0] = true // 空字符串可以被表示出来

	for j := 0; j <= len(s); j++ {
		for i := 0; i < len(wordDict); i++ {
			word := wordDict[i]
			l := len(word)

			// 当前位置 j 是否可以从 j-l 长度切出一个 word
			if j >= l && dp[j-l] && s[j-l:j] == word {
				dp[j] = true
				break // 已经能拼出 j 了，跳过后面判断
			}
		}
	}

	return dp[len(s)]
}

func wordBreak_wrong(s string, wordDict []string) bool {
	// dp[j] 表示：s[0:j]（前 j 个字符）是否可以被 wordDict 中的单词拼出
	dp := make([]bool, len(s)+1)
	dp[0] = true // 空字符串可以被表示出来

	// 反过来一样正确？ 🙅！！！是错误的
	for _, word := range wordDict {
		l := len(word)
		for j := l; j <= len(s); j++ {
			if dp[j-l] && s[j-l:j] == word {
				dp[j] = true
			}
		}
	}

	return dp[len(s)]
}
