package day01

import "testing"

func TestFindDuplicateFrequency(t *testing.T) {
	var input = `+1
-3
+1
+2
-4`

	var list = make(map[int64]bool)
	if x := FindDuplicateFrequency(input, 0, list); x != 1 {
		t.Errorf("Falsch, der Wert war %d", x)
	}
}

func TestFindDuplicateFrequency_Second_Run(t *testing.T) {
	var input = `+8
-4
-6
+2`

	var list = make(map[int64]bool)
	if x := FindDuplicateFrequency(input, 0, list); x != 8 {
		t.Errorf("Falsch, der Wert war %d", x)
	}
}

func TestFindDuplicateFrequency_New_Line_At_End(t *testing.T) {
	var input = `+8
-4
-6
+2
`

	var list = make(map[int64]bool)
	if x := FindDuplicateFrequency(input, 0, list); x != 8 {
		t.Errorf("Falsch, der Wert war %d", x)
	}
}
