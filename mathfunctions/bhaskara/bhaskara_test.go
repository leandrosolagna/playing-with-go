package bhaskara

import (
	"testing"
)

func TestDelta(t *testing.T) {
	expected := 196
	result := GetDelta(1, 12, -13)
	if result != expected {
		t.Errorf("The result is incorrect, got: %d, instead of %d", result, expected)
	}
}

func TestRoot(t *testing.T) {
	expected := 14
	result := GetRoot(196)
	if result != expected {
		t.Errorf("The result is incorrect, got: %d, instead of %d", result, expected)
	}
}

func TestBhaskaraMethod(t *testing.T) {
	x1_test, x2_test := BhaskaraMethod(1, 12, 14)
	if x1_test != 1 || x2_test != -13 {
		t.Errorf("Test case 1 failed: expected (%d, %d), got (%d, %d)", 1, -13, x1_test, x2_test)
	}
}
