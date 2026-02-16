package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/takymt/memo/internal/memo"
)

func captureStdout(t *testing.T, fn func() error) (string, error) {
	t.Helper()

	origStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return "", err
	}
	os.Stdout = w

	runErr := fn()

	_ = w.Close()
	os.Stdout = origStdout

	body, readErr := io.ReadAll(r)
	if readErr != nil {
		return "", readErr
	}
	return string(body), runErr
}

func TestRunListTodayAndWeek(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(t.TempDir(), "config.json")
	if err := memo.SaveConfig(configPath, memo.Config{MemoDir: tempDir}); err != nil {
		t.Fatal(err)
	}

	today := time.Now().Format("20060102")
	threeDaysAgo := time.Now().AddDate(0, 0, -3).Format("20060102")
	eightDaysAgo := time.Now().AddDate(0, 0, -8).Format("20060102")

	files := []string{
		today + "_today.md",
		threeDaysAgo + "_week.md",
		eightDaysAgo + "_old.md",
	}
	for _, name := range files {
		if err := os.WriteFile(filepath.Join(tempDir, name), []byte("# test\n"), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	todayOut, err := captureStdout(t, func() error {
		return runList(configPath, []string{"--today"})
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(todayOut, files[0]) {
		t.Fatalf("expected today file, got: %s", todayOut)
	}
	if strings.Contains(todayOut, files[1]) || strings.Contains(todayOut, files[2]) {
		t.Fatalf("unexpected file in today output: %s", todayOut)
	}

	weekOut, err := captureStdout(t, func() error {
		return runList(configPath, []string{"--week"})
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(weekOut, files[0]) || !strings.Contains(weekOut, files[1]) {
		t.Fatalf("expected this week's files, got: %s", weekOut)
	}
	if strings.Contains(weekOut, files[2]) {
		t.Fatalf("unexpected old file in week output: %s", weekOut)
	}
}
