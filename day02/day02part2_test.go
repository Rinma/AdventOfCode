package day02

import "testing"

func TestFindSimilar(t *testing.T) {
	var expected = "abcefgh"
	var input = `abcdefgh
abciefgh`

	if x := FindSimilar(input); x != expected {
		t.Errorf("The result should be %s but was %s", expected, x)
	}
}

func TestFindSimilar_Two(t *testing.T) {
	var expected = "abcefgh"
	var input = `abcdefgh
egabreds
abciefgh
trasvwsd`

	if x := FindSimilar(input); x != expected {
		t.Errorf("The result should be %s but was %s", expected, x)
	}
}
