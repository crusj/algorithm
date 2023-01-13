package leetcode

// [1 0 0 0 0 0 1]
func canPlaceFlowers(flowerbed []int, n int) bool {
	i, j := 0, 0
	ans := 0
	for i < len(flowerbed) {
		for i < len(flowerbed) && flowerbed[i] != 0 {
			i++
		}
		if i == len(flowerbed) {
			break
		}

		j = i + 1
		for j < len(flowerbed) && flowerbed[j] == 0 {
			j++
		}

		ans += (j - i - 1) / 2
		p := (j - i) % 2
		if p == 0 && (i-1 < 0 || j >= len(flowerbed)) {
			ans++
		}
		if p == 1 && i-1 < 0 && j >= len(flowerbed) {
			ans++
		}

		i = j + 1
	}

	return ans >= n
}
