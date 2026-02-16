package memo

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func SearchByFileName(dir, query string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}

	needle := strings.ToLower(strings.TrimSpace(query))
	if needle == "" {
		return []string{}, nil
	}

	var result []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if !strings.HasSuffix(name, ".md") {
			continue
		}
		if fuzzyContains(strings.ToLower(name), needle) {
			result = append(result, filepath.Join(dir, name))
		}
	}

	sort.Strings(result)
	return result, nil
}

func fuzzyContains(text, pattern string) bool {
	j := 0
	for i := 0; i < len(text) && j < len(pattern); i++ {
		if text[i] == pattern[j] {
			j++
		}
	}
	return j == len(pattern)
}
