package leetcode

/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
说明：你不能倾斜容器。

*/
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		v := (j - i) * min(height[i], height[j])
		if v > max {
			max = v
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return max
}
func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
