package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("x")
	expected := "xxxxx"

	if repeated != expected {
		t.Errorf("expected %v, but got %v", expected, repeated)
	}
}