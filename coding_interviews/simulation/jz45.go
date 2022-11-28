package simulation

import (
	"strconv"
	"strings"
)

func PrintMinNumber(numbers []int) string {
	numbersStr := make([]string, len(numbers))
	for i, v := range numbers {
		numbersStr[i] = strconv.Itoa(v)
	}

	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			a := numbersStr[j] + numbersStr[j+1]
			b := numbersStr[j+1] + numbersStr[j]
			if a > b {
				numbersStr[j], numbersStr[j+1] = numbersStr[j+1], numbersStr[j]
			}
		}
	}

	return strings.Join(numbersStr, "")
}
