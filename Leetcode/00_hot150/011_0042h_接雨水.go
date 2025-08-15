// https://leetcode.cn/problems/trapping-rain-water/description/?envType=study-plan-v2&envId=top-interview-150

package hot150

// trap 计算给定高度数组中可接的雨水总量
func trap(height []int) (res int) {
	n := len(height)
	l := 0     // 左边界索引
	store := 0 // 临时累加的雨水量

	// 第一次从左向右扫描，遇到更高的左边界时就收集雨水
	for i := 1; i < n; i++ {
		if height[i] >= height[l] {
			// 当前高度比左边界高或相等，说明可以结算之前积累的雨水
			l = i        // 更新左边界
			res += store // 加入当前积水
			store = 0    // 重置积水计数器
		} else {
			// 当前高度低于左边界，可以积水，积水量为差值
			store += height[l] - height[i]
		}
	}

	// 第二次从右向左扫描，处理右侧未被左边界捕捉到的雨水
	l, store = n-1, 0 // 重置右边界索引和积水量
	for i := n - 2; i >= 0; i-- {
		if height[i] > height[l] {
			// 注意这里是 > 而不是 >=，防止与前一遍重复计算
			l = i
			res += store
			store = 0
		} else {
			store += height[l] - height[i]
		}
	}
	return
}
