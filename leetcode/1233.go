package leetcode

/*
输入：folder = []
输出：["/a","/c/d","/c/f"]
解释："/a/b" 是 "/a" 的子文件夹，而 "/c/d/e" 是 "/c/d" 的子文件夹。
*/
func removeSubfolders(folder []string) []string {
	var container [101]map[string]struct{}
	indexs := make(map[string][]int, 0)
	for _, v := range folder {
		times := 0
		for i, ch := range v {
			if ch == '/' {
				times++
				indexs[v] = append(indexs[v], i)
			}
		}

		// 初始化
		if container[times] == nil {
			container[times] = map[string]struct{}{}
		}
		container[times][v] = struct{}{}
	}

	ans := make([]string, 0)
	for i := 100; i >= 1; i-- {
		if len(container[i]) == 0 {
			continue
		}

		tmp := make([]string, 0, len(ans)+len(container[i]))
		for _, str := range ans {
			tstr := str[:indexs[str][i]]
			if _, exists := container[i][tstr]; !exists {
				tmp = append(tmp, str)
			}
		}

		for str := range container[i] {
			tmp = append(tmp, str)
		}
		ans = tmp
	}

	return ans
}
