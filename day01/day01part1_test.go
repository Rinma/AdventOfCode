package day01_test

import (
	"AdventOfCode/day01"
	"testing"
)

func TestCalculate(t *testing.T) {
	var input = `+1`

	if x := day01.CalculateFinalFrequency(input); x != 1 {
		t.Errorf("Falsch, der Wert war %d", x)
	}
}

func TestCalculate_Two_Numbers(t *testing.T) {
	var input = `+1
-2`

	if x := day01.CalculateFinalFrequency(input); x != -1 {
		t.Errorf("Falsch, der Wert war %d", x)
	}
}

func TestCalculate_Three_Numbers(t *testing.T) {
	var input = `+1
-2
+3`

	if x := day01.CalculateFinalFrequency(input); x != 2 {
		t.Errorf("Falsch, der Wert war %d", x)
	}
}

func TestCalculate_New_Line_At_End(t *testing.T) {
	var input = `+1
-2
+3
`

	if x := day01.CalculateFinalFrequency(input); x != 2 {
		t.Errorf("Falsch, der Wert war %d", x)
	}
}
