//https://leetcode.cn/problems/jump-game-ii/description/
// m

package greedy

func jump(nums []int) (steps int) {
	var curMaxDis int // 当前能跳到的最远距离（当前层的最远边界）
	var nextDis int   // 下一层能跳到的最远距离

	for i := 0; i < len(nums); i++ {
		// 如果当前能到达的最远距离已经到达或超过终点，直接跳出
		if curMaxDis >= len(nums)-1 {
			break
		}

		// 更新从当前位置出发所能到达的最远位置
		nextDis = max(nextDis, i+nums[i])

		// 如果走到当前层的边界，准备跳到下一层
		if i == curMaxDis {
			curMaxDis = nextDis // 切换到下一层的边界
			steps++             // 计数跳跃次数
		}
	}
	return
}
