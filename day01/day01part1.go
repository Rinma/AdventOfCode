package day01

import (
	"strconv"
	"strings"
)

func CalculateFinalFrequency(input string) int64 {
	var ia = strings.Split(input, "\n")
	var result int64 = 0

	for _, element := range ia {
		if len(element) > 0 {
			operation := element[0]
			var number, _ = strconv.ParseInt(element[1:], 0, 0)
			if operation == '+' {
				result += number
			} else {
				result -= number
			}
		}
	}

	return result
}
