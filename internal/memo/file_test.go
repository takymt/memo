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

func TestFileNameFromDescriptionJapanese(t *testing.T) {
	got := FileNameFromDescription("買い物 メモ")
	if !strings.HasSuffix(got, "_買い物_メモ.md") {
		t.Fatalf("unexpected filename: %s", got)
	}
}
