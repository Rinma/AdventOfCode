package day01

import (
	"strconv"
	"strings"
)

func FindDuplicateFrequency(input string, result int64, list map[int64]bool) int64 {
	var ia = strings.Split(input, "\n")

	var foundDuplicate = false

	for _, element := range ia {
		if len(element) > 0 {
			operation := element[0]
			var number, _ = strconv.ParseInt(element[1:], 0, 0)
			if operation == '+' {
				result += number
			} else {
				result -= number
			}

			if list[result] == false {
				list[result] = true
			} else {
				foundDuplicate = true
				return result
			}
		}
	}

	if !foundDuplicate {
		return FindDuplicateFrequency(input, result, list)
	}

	return 0
}
