package mymath

import "testing"

func TestSub(t *testing.T) {
	got := Sub(1, 2)
	want := -1
	if got != want {
		t.Errorf("Sub(1, 2) = %d, want %d", got, want)
	}
}
