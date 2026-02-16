package memo

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode/utf8"
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
	textRunes := []rune(text)
	patternRunes := []rune(pattern)
	if len(patternRunes) == 0 {
		return true
	}
	if !utf8.ValidString(text) || !utf8.ValidString(pattern) {
		return false
	}

	j := 0
	for i := 0; i < len(textRunes) && j < len(patternRunes); i++ {
		if textRunes[i] == patternRunes[j] {
			j++
		}
	}
	return j == len(patternRunes)
}
