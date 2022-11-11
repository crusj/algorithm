package dp

// 有一种将字母编码成数字的方式：'a'->1, 'b->2', ... , 'z->26'。
// 现在给一串数字，返回有多少种可能的译码结果

/* example1
输入： "12"
返回值： 2
说明： 2种可能的译码结果（”ab” 或”l”）
*/

/* example2
输入： "31717126241541717"
返回值： 192
说明： 192种可能的译码结果
*/

// dp[n]是字符串前n个字符的译码结果
// base
// dp[0]  = 1  0为索引下标
// dp[n] = (dp[n - 1]) || (dp[n - 1] + dp[n - 2]  // n - 2 >= 0 and nums[n-2:n] <= 26)
func solve(nums string) int {
	numsByte := []byte(nums)
	if (len(nums) == 0) || len(nums) >= 1 && numsByte[0] == '0' {
		return 0
	}

	var dp [90]int
	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		curByte, preByte := numsByte[i], numsByte[i-1]
		a, b := curByte == '0', (preByte > '2' || preByte == '0') || (preByte == '2' && curByte > '6')
		switch {
		// 1、2位字符都无效
		case a && b:
			return 0
		// 2位数字有效
		case a && !b:
			if i == 1 {
				dp[1] = 1
				continue
			}

			dp[i] = dp[i-2]
			// 1位数字有效
		case !a && b:
			if i == 1 {
				dp[1] = 1
				continue
			}

			dp[i] = dp[i-1]
			// 1、2位数字都有效
		default:
			if i == 1 {
				dp[1] = 2
				continue
			}
			dp[i] = dp[i-1] + dp[i-2]
		}

	}

	return dp[len(nums)-1]
}
