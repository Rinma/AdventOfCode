package day02

import "testing"

func TestCalculatePackageChecksum(t *testing.T) {
	var expected = 1
	var input = `ababcbe`

	if x := CalculatePackageChecksum(input); x != expected {
		t.Errorf("Result is wrong, expected %d it is: %d", expected, x)
	}
}

func TestCalculatePackageChecksum_Two(t *testing.T) {
	var expected = 2
	var input = `ababcbe
aaebcdf`
	if x := CalculatePackageChecksum(input); x != expected {
		t.Errorf("Result is wrong, expected %d it is: %d", expected, x)
	}
}

func TestCalculatePackageChecksum_Three(t *testing.T) {
	var expected = 6
	var input = `ababcbe
aaebcdf
aeaefec`
	if x := CalculatePackageChecksum(input); x != expected {
		t.Errorf("Result is wrong, expected %d it is: %d", expected, x)
	}
}

func TestCalculatePackageChecksum_Example(t *testing.T) {
	var expected = 12
	var input = `abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`
	if x := CalculatePackageChecksum(input); x != expected {
		t.Errorf("Result is wrong, expected %d it is: %d", expected, x)
	}
}
