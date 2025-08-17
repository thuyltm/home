package fortune

import (
	"slices"
	"testing"
)

func TestGet(t *testing.T) {
	msg := Get()
	if i := slices.Index(fortunes, msg); i < 0 {
		t.Error("Expected a fortune, got an empty string")
	}
}
