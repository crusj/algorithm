package re

import "fmt"

func tower(n int, a, b, c string) {
	if n == 1 {
		fmt.Println("将" + a + "移动到" + c)
	} else {
		tower(n-1, a, c, b)
		fmt.Println("将" + a + "移动到" + c)
		tower(n-1, b, a, c)
	}
}
