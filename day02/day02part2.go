package day02

import (
	"strings"
)

func FindSimilar(input string) string {
	var i = strings.Split(input, "\n")
	var result = ""

	for outterIndex := 0; outterIndex < len(i)-1; outterIndex++ {
		for index := 0; index < len(i)-1; index++ {
			result = compareTwoStrings(i[outterIndex], i[index+1])

			if result != "" {
				return result
			}
		}
	}

	return result
}

func compareTwoStrings(input1 string, input2 string) string {
	different := 0
	var r = ""

	for i, c := range input1 {
		var l = rune(input2[i])
		if c != l {
			different++
		} else {
			r += string(l)
		}
	}

	if different == 1 {
		return r
	}

	return ""
}
