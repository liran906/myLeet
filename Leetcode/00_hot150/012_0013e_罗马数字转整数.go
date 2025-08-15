// https://leetcode.cn/problems/roman-to-integer/description/?envType=study-plan-v2&envId=top-interview-150

package hot150

// romanToInt 将罗马数字字符串 s 转换为整数
func romanToInt(s string) (res int) {
	// 定义罗马字符与对应整数值的映射
	mps := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// 遍历字符串中的每个字符
	for i := range s {
		ch := s[i]

		// 如果当前字符比前一个字符代表的值大，说明前一个字符是“减法组合”
		// 例如：IV = 5 - 1 = 4，需要减去两倍前一个值（因为之前加过一次）
		if i > 0 && mps[ch] > mps[s[i-1]] {
			res -= mps[s[i-1]] * 2
		}

		// 加上当前字符的值（无论是否是减法组合）
		res += mps[ch]
	}

	return
}
