// https://leetcode.cn/problems/combination-sum-iv/

package dp

// 为啥一维 dp 的时候遍历顺序影响是排列还是组合，我也没特别搞懂，但只能说：

/*
https://programmercarl.com/0518.%E9%9B%B6%E9%92%B1%E5%85%91%E6%8D%A2II.html#%E4%B8%80%E7%BB%B4dp%E8%AE%B2%E8%A7%A3

我们先来看 外层for循环遍历物品（钱币），内层for遍历背包（金钱总额）的情况。
for (int i = 0; i < coins.size(); i++) { // 遍历物品
    for (int j = coins[i]; j <= amount; j++) { // 遍历背包容量
        dp[j] += dp[j - coins[i]];
    }
}
假设：coins[0] = 1，coins[1] = 5。
那么就是先把1加入计算，然后再把5加入计算，得到的方法数量只有{1, 5}这种情况。而不会出现{5, 1}的情况。
所以这种遍历顺序中dp[j]里计算的是组合数！


如果把两个for交换顺序，代码如下：
for (int j = 0; j <= amount; j++) { // 遍历背包容量
    for (int i = 0; i < coins.size(); i++) { // 遍历物品
        if (j - coins[i] >= 0) dp[j] += dp[j - coins[i]];
    }
}
背包容量的每一个值，都是经过 1 和 5 的计算，包含了{1, 5} 和 {5, 1}两种情况。
此时dp[j]里算出来的就是排列数！
*/

// combinationSum4 计算能凑出 target 的所有排列数（顺序不同算不同组合）
// 输入：nums - 可用的数字列表（每个数字可重复使用），target - 目标和
// 输出：所有排列的个数（顺序不同算不同）
func combinationSum4(nums []int, target int) int {
	// dp[j] 表示：凑出总和为 j 的**排列数**
	dp := make([]int, target+1)

	// 初始化：dp[0] = 1 表示总和为 0 的排列只有一种（什么都不选）
	dp[0] = 1

	// 遍历金额（外层），表示我们依次构建每个金额 j
	for j := 0; j <= target; j++ {
		// 对于每一个金额 j，我们尝试用每一个 nums[i] 来“跳”到 j
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				// 说明当前硬币可以参与构成 j
				// dp[j] += dp[j - nums[i]] 表示：
				// 所有能凑成 j-nums[i] 的排列，只需再加上 nums[i]，就能构成 j
				dp[j] += dp[j-nums[i]]
			}
		}
	}

	// dp[target] 即为凑出 target 的所有排列数
	return dp[target]
}
