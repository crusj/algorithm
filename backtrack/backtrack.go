package backtrack

import (
	"fmt"
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
