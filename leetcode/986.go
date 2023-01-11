package leetcode

/*
输入：
firstList = [[0,2],[5,10],[13,23],[24,25]],
secondList = [[1,5],[8,12],[15,24],[25,26]]
输出：[[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
*/
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	i, j := 0, 0
	ret := make([][]int, 0)
	for {
		if i == len(firstList) || j == len(secondList) {
			break
		}

		if firstList[i][1] > secondList[j][1] {
			if firstList[i][0] > secondList[j][1] {
				j++
				continue
			}

			if firstList[i][0] <= secondList[j][0] {
				ret = append(ret, []int{secondList[j][0], secondList[j][1]})
			} else {
				ret = append(ret, []int{firstList[i][0], secondList[j][1]})
			}
			j++
		} else if firstList[i][1] == secondList[j][1] {
			start := firstList[i][0]
			if secondList[j][0] > start {
				start = secondList[j][0]
			}
			ret = append(ret, []int{start, firstList[i][1]})
			i++
			j++
		} else {
			// first[i][1] < second[j][1]
			if secondList[j][0] > firstList[i][1] {
				i++
				continue
			}

			if secondList[j][0] <= firstList[i][0] {
				ret = append(ret, []int{firstList[i][0], firstList[i][1]})
			} else {
				ret = append(ret, []int{secondList[j][0], firstList[i][1]})
			}
			i++
		}

	}

	return ret
}
