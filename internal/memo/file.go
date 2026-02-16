package memo

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"
)

var separatorPattern = regexp.MustCompile(`[\s　]+`)

func FileNameFromDescription(description string) string {
	d := strings.TrimSpace(description)
	d = separatorPattern.ReplaceAllString(d, "_")
	d = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-' {
			return r
		}
		return -1
	}, d)
	if d == "" {
		d = "memo"
	}
	return fmt.Sprintf("%s_%s.md", time.Now().Format("20060102"), d)
}
