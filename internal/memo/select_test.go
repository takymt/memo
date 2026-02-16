package memo

import "testing"

func TestBestMatch(t *testing.T) {
	got := BestMatch([]string{"b.md", "a.md"})
	if got != "b.md" {
		t.Fatalf("unexpected best match: %s", got)
	}
}
