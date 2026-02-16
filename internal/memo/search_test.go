package memo

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSearchByFileName(t *testing.T) {
	tempDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(tempDir, "20260101_hello.md"), []byte("# hello\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tempDir, "20260101_world.md"), []byte("# world\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	matches, err := SearchByFileName(tempDir, "hlo")
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 1 {
		t.Fatalf("expected 1 match, got %d", len(matches))
	}
}
