package search

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param str string字符串
 * @return string字符串一维数组
 */
func Permutation(str string) []string {
	if len(str) == 0 {
		return nil
	}

	// write code here
	alreadyExists := make(map[string]struct{})
	ret := make([]string, 0)

	bys := []byte(str)
	var tf func(visited []bool, path []byte)
	tf = func(visited []bool, path []byte) {
		if len(path) == len(str) {
			if _, exists := alreadyExists[string(path)]; !exists {
				alreadyExists[string(path)] = struct{}{}
				ret = append(ret, string(path))
			}

			return
		}

		for i, v := range bys {
			if visited[i] {
				continue
			}

			visited[i] = true
			path = append(path, v)

			tf(visited, path)
			path = path[0 : len(path)-1]
			visited[i] = false
		}
	}

	tf(make([]bool, len(str)), make([]byte, 0))

	return ret
}
