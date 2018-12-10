package day02

import "strings"

func hasAmountOfLetters(word string, amount int) bool {
	var letters = make(map[int32]int)

	for _, c := range word {
		letters[c] += 1
	}

	for _, j := range letters {
		if j == amount {
			return true
		}
	}

	return false
}

func CalculatePackageChecksum(input string) int {
	var i = strings.Split(input, "\n")
	var resultDouble = 0
	var resultTriple = 0

	for _, e := range i {
		if hasAmountOfLetters(e, 2) {
			resultDouble++
		}

		if hasAmountOfLetters(e, 3) {
			resultTriple++
		}
	}

	return resultDouble * resultTriple
}
