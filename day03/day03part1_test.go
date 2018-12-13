package day03

import "testing"

func TestCalculateOverlappingAreas(t *testing.T) {
	var expected = 0
	var in = `#1 @ 1,3: 4x4`

	if area := CalculateOverlappingAreas(in); area != expected {
		t.Fatalf(`Expected %d but got %d`, expected, area)
	}
}

func TestCalculateOverlappingAreas_Two(t *testing.T) {
	var expected = 4
	var in = `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4`

	if area := CalculateOverlappingAreas(in); area != expected {
		t.Fatalf(`Expected %d but got %d`, expected, area)
	}
}

func TestCalculateOverlappingAreas_Three(t *testing.T) {
	var expected = 4
	var in = `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`

	if area := CalculateOverlappingAreas(in); area != expected {
		t.Fatalf(`Expected %d but got %d`, expected, area)
	}
}

func TestCalculateOverlappingAreas_Four(t *testing.T) {
	var expected = 20
	var in = `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 3,1: 4x4`

	if area := CalculateOverlappingAreas(in); area != expected {
		t.Fatalf(`Expected %d but got %d`, expected, area)
	}
}

func TestCalculateOverlappingAreas_Five(t *testing.T) {
	var expected = 11
	var in = `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2
#4 @ 1,3: 3x3`

	if area := CalculateOverlappingAreas(in); area != expected {
		t.Fatalf(`Expected %d but got %d`, expected, area)
	}
}

func TestCalculateOverlappingAreas_Six(t *testing.T) {
	var expected = 400
	var in = `#1 @ 10,30: 40x40
#2 @ 30,10: 40x40
#3 @ 50,50: 20x20`

	if area := CalculateOverlappingAreas(in); area != expected {
		t.Fatalf(`Expected %d but got %d`, expected, area)
	}
}
