package search

import (
	"strconv"
)

/**
* 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
*
*
* @param n int整型
* @return int整型
01234567891011121314
* [0,10) [10,100), [100,1000) [1000,10000)
*/
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	i := 1
	a, b := 0, 10
	for n > (b-a)*i {
		t := (b - a) * i
		n -= t
		i++
		a, b = b, b*10
	}

	num := n/i + a
	index := n % i
	x := strconv.Itoa(num)
	return int([]byte(x)[index] - '0')
}
