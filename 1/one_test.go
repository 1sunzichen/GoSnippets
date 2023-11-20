package One

import "testing"

func TestMain(t *testing.T) {
	a := 2
	b := 3
	want := 3
	got := If(a > b, a, b)
	if got != want {
		t.Errorf("got %d,want %d", got, want)
	}
}
