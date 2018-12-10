package main

import (
	"AdventOfCode/day01"
	"AdventOfCode/day02"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// DAY 01
	day01input, err := ioutil.ReadFile("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	day01part1(string(day01input))
	day01part2(string(day01input))

	// DAY 02
	day02input, err := ioutil.ReadFile("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	day02part1(string(day02input))
	day02part2(string(day02input))
}

func day01part1(input string) {
	result := day01.CalculateFinalFrequency(input)

	fmt.Printf("Result of Day01 part 1 is: %d\n", result)
}

func day01part2(input string) {
	var list = make(map[int64]bool)
	result := day01.FindDuplicateFrequency(input, 0, list)

	fmt.Printf("Result of Day01 part 2 is: %d\n", result)
}

func day02part1(input string) {
	result := day02.CalculatePackageChecksum(input)

	fmt.Printf("Result of Day02 part 1 is: %d\n", result)
}

func day02part2(input string) {
	result := day02.FindSimilar(input)

	fmt.Printf("Result of Day02 part 2 is: %s\n", result)
}
