package memo

import (
	"strings"
	"testing"
)

func TestFileNameFromDescription(t *testing.T) {
	got := FileNameFromDescription("hello world")
	if !strings.HasSuffix(got, "_hello_world.md") {
		t.Fatalf("unexpected filename: %s", got)
	}
}
