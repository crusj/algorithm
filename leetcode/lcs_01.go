package leetcode

/*
小扣打算给自己的VScode安装使用插件，初始状态下带宽每分钟可以完成1个插件的下载。假定每分钟选择以下两种策略之一:
使用当前带宽下载插件
将带宽加倍（下载插件数量随之加倍）
请返回小扣完成下载n个插件最少需要多少分钟。
注意：实际的下载的插件数量可以超过n个

初始贷款每分钟完成一个插件下载
选择:
使用当前宽带下载。
将代管加倍，下载插件的数量也会翻倍。
*/

// a 2a => 选择需要2分,加倍需要先加倍在选择也是需要两分。但加倍更利于下一步。
func leastMinutes(n int) int {
	minutes := 0
	speed := 1
	for n > 0 {
		if n <= speed {
			n -= speed
			minutes++
		} else {
			minutes++
			speed *= 2
		}
	}

	return minutes
}

// 回溯超时
func leastMinutes2(n int) int {
	var fn func(speed, now, total, minutes int) int
	fn = func(speed, now, total, minutes int) int {
		if now >= total {
			return minutes
		}
		if speed >= total-now {
			return minutes + 1
		}

		a1 := fn(speed, now+speed, total, minutes+1)
		a2 := fn(2*speed, now, total, minutes+1)
		if a1 < a2 {
			return a1
		}

		return a2
	}

	return fn(1, 0, n, 0)
}
