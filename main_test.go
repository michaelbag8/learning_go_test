package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := multiplyNumber(2, 3)
	expected := 6
	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func TestAddNumber(t *testing.T) {
	want := AddNumber(3, 10)
	got := 13
	if want != got {
		t.Errorf("AddNumber(3,10) is %d not %d", want, got)
	}
}

// Table-driven tests
func TestAdds(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{2, 3, 5},
		{3, 15, 18},
		{4, 10, 14},
		{0, 0, 0},
		{-1, -1, -2},
	}
	for _, test := range tests {
		got := AddNumber(test.a, test.b)
		if got != test.want {
			t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, got, test.want)
		}

	}
}
