package structure

import (
	"reflect"
)

func IsPopOrder(pushV []int, popV []int) bool {
	var tf func(pushV []int, stack []int, path []int) bool
	tf = func(pushV, stack []int, path []int) bool {
		if len(path) == len(popV) {
			return reflect.DeepEqual(path, popV)
		} else if len(path) > 0 {
			if !reflect.DeepEqual(path, popV[0:len(path)]) {
				return false
			}
		}

		// 出栈
		if len(stack) > 0 {
			if tf(pushV, stack[0:len(stack)-1], append(path, stack[len(stack)-1])) {
				return true
			}
		}

		// 入栈
		if len(pushV) > 0 {
			return tf(pushV[1:], append(stack, pushV[0]), path)
		}

		return false
	}

	return tf(pushV, []int{}, []int{})
}
