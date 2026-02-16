package memo

import (
	"path/filepath"
	"testing"
)

func TestDefaultMemoDirUsesXDGDataHome(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", "/tmp/xdg-data")

	got, err := DefaultMemoDir()
	if err != nil {
		t.Fatal(err)
	}

	want := filepath.Join("/tmp/xdg-data", "memo")
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestLoadOrDefaultConfigFallsBackToDefaultMemoDir(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", "/tmp/xdg-data")

	cfg, err := LoadOrDefaultConfig(filepath.Join(t.TempDir(), "missing-config.json"))
	if err != nil {
		t.Fatal(err)
	}

	want := filepath.Join("/tmp/xdg-data", "memo")
	if cfg.MemoDir != want {
		t.Fatalf("got %q, want %q", cfg.MemoDir, want)
	}
}
