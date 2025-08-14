// https://leetcode.cn/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=top-interview-150

package hot150

// 看官方解析
func productExceptSelf(nums []int) []int {
	n := len(nums)        // 获取输入数组的长度
	ans := make([]int, n) // 创建返回结果数组（与输入等长）

	// 【第一遍遍历】从左到右构建“左侧乘积”
	ans[0] = 1 // ans[i] 初始表示 nums[0:i] 的乘积（不含 nums[i]）
	for i := 1; i < n; i++ {
		// ans[i] = 左边所有数的乘积 = ans[i-1] * nums[i-1]
		ans[i] = ans[i-1] * nums[i-1]
	}

	// 【第二遍遍历】从右到左乘上“右侧乘积”
	suffix := 1 // suffix 表示 nums[i+1:] 的乘积
	for i := n - 1; i >= 0; i-- {
		ans[i] *= suffix  // 把右侧乘积乘进来
		suffix *= nums[i] // 更新 suffix，准备下一个 i 的右侧乘积
	}

	return ans
}
