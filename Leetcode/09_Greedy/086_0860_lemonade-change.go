// https://leetcode.cn/problems/lemonade-change/
// e

package greedy

func lemonadeChange_(bills []int) bool {
	// 初始化手头零钱的数量：分别记录 5元、10元、20元的张数
	m := map[int]int{5: 0, 10: 0, 20: 0}

	for i := range bills {
		switch bills[i] {
		case 5:
			// 顾客付5元，不用找零，直接收下
			m[5]++

		case 10:
			// 顾客付10元，需要找回5元
			m[5]--
			if m[5] < 0 {
				// 如果没有5元纸币，无法找零
				return false
			}
			m[10]++

		case 20:
			// 顾客付20元，需要找回15元
			if m[10] > 0 && m[5] > 0 {
				// 优先使用一张10元和一张5元（最优策略）
				m[10]--
				m[5]--
			} else if m[5] >= 3 {
				// 否则只能用三张5元找零
				m[5] -= 3
			} else {
				// 两种方案都失败，无法找零
				return false
			}
			m[20]++ // 实际这里是否加 20 记录都无所谓
		}
	}

	return true // 全部顾客都成功找零
}

// 优化写法
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		} else {
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
