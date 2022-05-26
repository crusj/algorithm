package backtrack

import (
	"fmt"
	"sort"
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
