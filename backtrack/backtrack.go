package backtrack

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// WholeArrangement 对一组数字进行全排列
func WholeArrangement(nums []int, backtrack []int) [][]int {
	results := make([][]int, 0)

	var fn func(nums []int, backtrack []int)

	fn = func(nums []int, backtrack []int) {
		if len(backtrack) == len(nums) {
			results = append(results, backtrack)
			return
		}

		for _, item := range nums {
			exists := false
			for _, tmp := range backtrack {
				if item == tmp {
					exists = true
					break
				}
			}

			if exists {
				continue
			}

			newBackTrack := make([]int, len(backtrack))
			copy(newBackTrack, backtrack)

			newBackTrack = append(newBackTrack, item)
			fn(nums, newBackTrack)
		}
	}

	fn(nums, backtrack)

	return results
}

// QueenN function    N皇后问题.
func QueenN(n int) {
	// init
	backtrack := make([][]int, n)
	for i := 0; i < n; i++ {
		backtrack[i] = make([]int, n)
	}

	total_nums := 0

	var fn func(n int, row int, backTrack [][]int)

	checkValid := func(row, col int, backtrack [][]int) bool {
		// check above
		for i := 0; i < row; i++ {
			if backtrack[i][col] != 0 {
				return false
			}
		}

		// check left above
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if backtrack[i][j] != 0 {
				return false
			}
		}

		// check right above
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if backtrack[i][j] != 0 {
				return false
			}
		}

		return true
	}

	fn = func(n int, row int, backTrack [][]int) {
		if row == n {
			total_nums++
			fmt.Println(total_nums)
			for _, row := range backTrack {
				for _, col := range row {
					fmt.Print(col)
				}
				fmt.Printf("\n")
			}
			fmt.Printf("\n\n")
			return
		}

		for i := 0; i < n; i++ {
			if !checkValid(row, i, backTrack) {
				continue
			}

			newBackTrack := make([][]int, n)
			for j := 0; j < n; j++ {
				newBackTrack[j] = make([]int, n)
				copy(newBackTrack[j], backTrack[j])
			}

			newBackTrack[row][i] = 1
			fn(n, row+1, newBackTrack)
		}
	}

	fn(n, 0, backtrack)
}

// Climb function   爬楼梯，一共有N层楼梯，有几种步数，求全部方案
func Climb(floor int, ways []int, path []int) {
	results := make([][]int, 0)

	var fn func(int, []int, []int)
	fn = func(floor int, ways []int, path []int) {
		if floor == 0 {
			results = append(results, path)

			return
		}

		for _, i := range ways {
			if floor-i < 0 {
				continue
			}

			newPath := make([]int, len(path))
			copy(newPath, path)
			newPath = append(newPath, i)

			fn(floor-i, ways, newPath)
		}
	}

	fn(floor, ways, []int{})

	fmt.Println(len(results))
	for _, item := range results {
		fmt.Println(item)
	}
}

// ClimbNums function    Climb Nums.
func ClimbNums(floor int) int {
	if floor == 0 {
		return 1
	} else if floor < 0 {
		return 0
	}

	return ClimbNums(floor-1) + ClimbNums(floor-2) + ClimbNums(floor-3)
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	digitsBytes := []byte(digits)
	m := map[byte][]byte{
		'1': {},
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	results := []string{}

	var fn func(row int, combine []byte)

	fn = func(offset int, combine []byte) {
		if len(combine) == len(digits) {
			results = append(results, string(combine))

			return
		}

		for _, c := range m[digitsBytes[byte(offset)]] {
			newCombie := make([]byte, len(combine))
			copy(newCombie, combine)
			newCombie = append(newCombie, c)
			fn(offset+1, newCombie)
		}
	}

	fn(0, []byte{})

	return results
}

// generateParenthesis function    leetcode 22
func generateParenthesis(n int) []string {
	results := []string{}

	checkValids := func(str []byte) bool {
		stack := make([]byte, 0, len(str))
		for _, v := range str {
			if v == '(' {
				stack = append(stack, v)
			} else {
				if len(stack) == 0 {
					return false
				}

				if stack[len(stack)-1] != '(' {
					return false
				}

				stack = stack[0 : len(stack)-1]
			}
		}

		return len(stack) == 0
	}

	var fn func(int, []byte)
	fn = func(i int, s []byte) {
		if len(s) == n*2 {
			if checkValids(s) {
				results = append(results, string(s))
			}

			return
		}

		newS := make([]byte, len(s))
		copy(newS, s)
		newS = append(newS, '(')
		fn(n-1, newS)

		newS = make([]byte, len(s))
		copy(newS, s)
		newS = append(newS, ')')
		fn(n-1, newS)
	}

	fn(n*2, []byte{})

	return results
}

// combinationSum function    leetcode39
func combinationSum(candidates []int, target int) [][]int {
	results := make([][]int, 0)
	exists := make(map[string]struct{})
	sort.Ints(candidates)
	var fn func(int, []int)
	fn = func(t int, path []int) {
		if t == 0 {
			sort.Ints(path)

			k := make([]byte, len(path))

			for i, c := range path {
				k[i] = byte(c)
			}

			if _, is := exists[string(k)]; is {
				return
			} else {
				results = append(results, path)
				exists[string(k)] = struct{}{}
			}

			return
		}

		for _, i := range candidates {
			if t-i < 0 {
				break
			}

			newPath := make([]int, len(path), len(path)+1)
			copy(newPath, path)
			newPath = append(newPath, i)

			fn(t-i, newPath)
		}
	}

	fn(target, []int{})

	return results
}

// combinationSum2 function    leetcode40
func combinationSum2(candidates []int, target int) [][]int {
	results := make([][]int, 0)
	sort.Ints(candidates)

	preSum := make([]int, len(candidates)+1) // 前缀和
	preSum[0] = 0
	for i := 0; i < len(candidates); i++ {
		preSum[i+1] = preSum[i] + candidates[i]
	}

	exists := make(map[string]struct{})
	times := make(map[int]int)
	for _, v := range candidates {
		if _, is := times[v]; is {
			times[v]++
		} else {
			times[v] = 1
		}
	}

	if len(times) == 1 {
		if target%candidates[0] == 0 {
			q := target / candidates[0]
			tmp := times[candidates[0]] - q
			if tmp >= 0 {
				s := make([]int, 0, q)
				for b := 0; b < q; b++ {
					s = append(s, candidates[0])
				}

				results = append(results, s)
				return results

			} else {
				return results
			}
		} else {
			return results
		}
	}

	var fn func(int, map[int]int)
	fn = func(t int, path map[int]int) {
		if t == 0 {
			tmp := make([]int, 0, len(path))
			for k, v := range path {
				for i := 1; i <= v; i++ {
					tmp = append(tmp, k)
				}
			}

			sort.Ints(tmp)
			k := make([]byte, len(tmp))

			for i, c := range tmp {
				k[i] = byte(c)
			}

			if _, is := exists[string(k)]; is {
				return
			} else {
				results = append(results, tmp)
				exists[string(k)] = struct{}{}
			}
		}

		newPath := make(map[int]int, len(path)+1)
		for k, v := range path {
			newPath[k] = v
		}

		for x, i := range candidates {
			if t-i < 0 {
				break
			}

			sum := preSum[len(candidates)] - preSum[x]
			if sum < t {
				break
			}

			if t, is := path[i]; is && t == times[i] {
				continue
			}

			if _, is := path[i]; is {
				newPath[i]++
			} else {
				newPath[i] = 1
			}

			fn(t-i, newPath)
			if newPath[i] > 1 {
				newPath[i]--
			} else {
				delete(newPath, i)
			}
		}
	}

	fn(target, make(map[int]int))

	return results
}

// exist function    leetcode79
func exist(board [][]byte, word string) bool {
	result := false
	wordBytes := []byte(word)

	isExists := func(row, col int, path []int) bool {
		for i := 0; i < len(path); i = i + 2 {
			if row == path[i] && col == path[i+1] {
				return true
			}
		}

		return false
	}

	var fn func(int, int, []int)
	fn = func(row, col int, path []int) {
		if len(path)/2 == len(wordBytes) {
			result = true
			return
		}

		// selects
		// top
		if row-1 >= 0 {
			if !isExists(row-1, col, path) { // must not repeat
				if board[row-1][col] == wordBytes[len(path)/2] { // valid
					newPath := make([]int, len(path), len(path)+2)
					copy(newPath, path)
					newPath = append(newPath, []int{row - 1, col}...)
					fn(row-1, col, newPath)
				}
			}
		}

		// bottom
		if row+1 < len(board) {
			if !isExists(row+1, col, path) { // must not repeat
				if board[row+1][col] == wordBytes[len(path)/2] { // valid
					newPath := make([]int, len(path), len(path)+2)
					copy(newPath, path)
					newPath = append(newPath, []int{row + 1, col}...)
					fn(row+1, col, newPath)
				}
			}
		}

		// left
		if col-1 >= 0 {
			if !isExists(row, col-1, path) { // must not repeat
				if board[row][col-1] == wordBytes[len(path)/2] { // valid
					newPath := make([]int, len(path), len(path)+2)
					copy(newPath, path)
					newPath = append(newPath, []int{row, col - 1}...)
					fn(row, col-1, newPath)
				}
			}
		}

		// right
		if col+1 < len(board[0]) {
			if !isExists(row, col+1, path) { // must not repeat
				if board[row][col+1] == wordBytes[len(path)/2] { // valid
					newPath := make([]int, len(path), len(path)+2)
					copy(newPath, path)
					newPath = append(newPath, []int{row, col + 1}...)
					fn(row, col+1, newPath)
				}
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != wordBytes[0] {
				continue
			}

			fn(i, j, []int{i, j})
		}
	}

	return result
}

// restoreIpAddresses function    剑指向Offer 087
func restoreIpAddresses(s string) []string {
	results := make([]string, 0)
	ls := len(s)

	var fn func(int, []string)
	fn = func(offset int, path []string) {
		if len(path) > 0 {
			if d, _ := strconv.Atoi(path[len(path)-1]); d > 255 {
				return
			}
		}

		if offset == ls {
			if len(path) == 4 {
				results = append(results, strings.Join(path, "."))
			}

			return
		} else {
			if len(path) >= 4 {
				return
			}
		}

		newPath := make([]string, len(path))
		copy(newPath, path)

		v := ls - offset

		if v >= 1 {
			newPath = append(newPath, s[offset:offset+1])
			fn(offset+1, newPath)
			newPath = newPath[0 : len(newPath)-1]
		}

		if s[offset:offset+1] == "0" {
			return
		}

		if v >= 2 {
			newPath = append(newPath, s[offset:offset+2])
			fn(offset+2, newPath)
			newPath = newPath[0 : len(newPath)-1]
		}

		if v >= 3 {
			newPath = append(newPath, s[offset:offset+3])
			fn(offset+3, newPath)
			newPath = newPath[0 : len(newPath)-1]
		}
	}

	fn(0, []string{})

	return results
}

// solveNQueens function    面试题 08.12. 八皇后
func solveNQueens(n int) [][]string {
	results := make([][]string, 0)

	// 初始化
	path := make([][]byte, n)
	for i := 0; i < n; i++ {
		path[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			path[i][j] = '.'
		}
	}

	isValid := func(row, col int, path [][]byte) bool {
		// top
		for i := row - 1; i >= 0; i-- {
			if path[i][col] == 'Q' {
				return false
			}
		}

		// left top
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if path[i][j] == 'Q' {
				return false
			}
		}

		// right top
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if path[i][j] == 'Q' {
				return false
			}
		}

		return true
	}

	var fn func(int, [][]byte)
	fn = func(row int, path [][]byte) {
		if row == n {
			tmp := make([]string, len(path))
			for k, v := range path {
				tmp[k] = string(v)
			}
			results = append(results, tmp)

			return
		}

		newPath := make([][]byte, n)
		for i := 0; i < len(path); i++ {
			newPath[i] = make([]byte, n)
			copy(newPath[i], path[i])
		}

		for col := 0; col < n; col++ {
			if isValid(row, col, path) {
				newPath[row][col] = 'Q'
				fn(row+1, newPath)
				newPath[row][col] = '.'
			}
		}
	}

	fn(0, path)

	return results
}

// exist function    剑指 Offer 12. 矩阵中的路径
func exist2(board [][]byte, word string) bool {
	is := false

	isValid := func(row, col int, h map[[2]int]struct{}) bool {
		if _, exists := h[[2]int{row, col}]; exists {
			return false
		}

		return board[row][col] == word[len(h)]
	}

	var fn func(row, col int, h map[[2]int]struct{}, path []int)
	fn = func(row, col int, h map[[2]int]struct{}, path []int) {
		if is {
			return
		}
		if len(path) == len(word)*2 {
			is = true
			return
		}

		newPath := make([]int, len(path))
		copy(newPath, path)

		newH := make(map[[2]int]struct{}, len(h))
		for k, v := range h {
			newH[k] = v
		}

		// top
		if row-1 >= 0 && isValid(row-1, col, newH) {
			newPath = append(newPath, row-1, col)
			t := [2]int{row - 1, col}
			newH[t] = struct{}{}
			fn(row-1, col, newH, newPath)
			newPath = newPath[0 : len(newPath)-2]
			delete(newH, t)
		}

		// down
		if row+1 < len(board) && isValid(row+1, col, newH) {
			newPath = append(newPath, row+1, col)
			t := [2]int{row + 1, col}
			newH[t] = struct{}{}
			fn(row+1, col, newH, newPath)
			newPath = newPath[0 : len(newPath)-2]
			delete(newH, t)
		}

		// left
		if col-1 >= 0 && isValid(row, col-1, newH) {
			newPath = append(newPath, row, col-1)
			t := [2]int{row, col - 1}
			newH[t] = struct{}{}
			fn(row, col-1, newH, newPath)
			newPath = newPath[0 : len(newPath)-2]
		}

		// right
		if col+1 < len(board[0]) && isValid(row, col+1, newH) {
			newPath = append(newPath, row, col+1)
			t := [2]int{row, col + 1}
			newH[t] = struct{}{}
			fn(row, col+1, newH, newPath)
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				fn(i, j, map[[2]int]struct{}{
					{i, j}: {},
				}, []int{i, j})
			}
		}
	}

	return is
}

// maxRectangle function    面试题 17.25. 单词矩阵
func maxRectangle(words []string) []string {
	results := make([]string, 0)

	isValid := func(s string, path []string) bool {
		if len(path) == 0 {
			return true
		}

		if len(s) != len(path[0]) {
			return false
		}

		for i := 0; i < len(path[0]); i++ {
			t := ""
			for j := 0; j < len(path); j++ {
				t = t + path[j][i:i+1]
			}

			t += s[i : i+1]
			exists := false
			for _, s := range words {
				if strings.HasPrefix(s, t) {
					exists = true
					break
				}
			}
			if !exists {
				return false
			}

		}

		return true
	}

	var fn func(path []string)
	fn = func(path []string) {
		if len(path)*len(path[0]) > len(results)*len(results[0]) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			results = tmp
		}

		newPath := make([]string, len(path))
		copy(newPath, path)

		for i := 0; i < len(words); i++ {
			if isValid(words[i], path) {
				newPath = append(newPath, words[i])
				fn(newPath)
				newPath = newPath[0 : len(newPath)-1] // back
			}
		}
	}
	for i := 0; i < len(words); i++ {
		if len(results) == 0 {
			results = []string{words[i]}
		}

		fn([]string{words[i]})
	}

	return results
}

// allPathsSourceTarget function    leetcode797
func allPathsSourceTarget(graph [][]int) [][]int {
	results := make([][]int, 0)

	var fn func(path []int)
	fn = func(path []int) {
		if path[len(path)-1] == len(graph)-1 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			results = append(results, tmp)

			return
		}

		newPath := make([]int, len(path))
		copy(newPath, path)

		d := graph[path[len(path)-1]]

		for i := 0; i < len(d); i++ {
			newPath = append(newPath, d[i])
			fn(newPath)
			newPath = newPath[0 : len(newPath)-1]
		}
	}

	for i := 0; i < len(graph[0]); i++ {
		fn([]int{0, graph[0][i]})
	}

	return results
}

// findTargetSumWays function    leetcode 494 目标和.
// todo go超时，可以看php版本，golang切片还要复制太麻烦.
func findTargetSumWays(nums []int, target int) int {
	result := 0

	var fn func(index int, path []int)
	fn = func(index int, path []int) {
		// 结束条件
		if index == len(nums) {
			sum := 0
			for i := 0; i < len(path); i++ {
				sum += path[i]
			}

			if sum == target {
				result++
			}

			return
		}

		// 复制路径
		newPath := make([]int, len(path))
		copy(newPath, path)

		// 选择加号
		fn(index+1, newPath)

		// 选择减号
		newPath[index] = -1 * newPath[index]
		fn(index+1, newPath)
	}

	fn(0, nums)

	return result
}

// findTargetSumWays function    leetcode 494 目标和.
// todo go超时，可以看php版本，golang切片还要复制太麻烦.
func findTargetSumWays2(nums []int, target int) int {
	result := 0

	var fn func(index int, sum int)
	fn = func(index int, sum int) {
		// 结束条件
		if index == len(nums) {
			if sum == target {
				result++
			}
			return
		}

		// 选择加号
		fn(index+1, sum+nums[index])
		// 选择减号
		fn(index+1, sum+nums[index]*-1)
	}

	fn(0, 0)

	return result
}
