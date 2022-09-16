package structure

import "strings"

func ReverseSentence(str string) string {
	if str == "" {
		return ""
	}

	stack := make([]string, 0)
	word := ""
	for i := 0; i <= len(str); i++ {
		if i == len(str) {
			stack = append(stack, word)
			break
		}

		if str[i:i+1] == " " {
			stack = append(stack, word)
			word = ""
			continue
		}

		word += str[i : i+1]
	}

	stackLen := len(stack)
	ret := ""
	for stackLen > 0 {
		ret += stack[stackLen-1] + " "
		stackLen--
	}
	// write code here
	return strings.TrimRight(ret, " ")
}
