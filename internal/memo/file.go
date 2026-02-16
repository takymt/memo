package memo

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var nonFileChar = regexp.MustCompile(`[^a-zA-Z0-9_-]+`)

func FileNameFromDescription(description string) string {
	d := strings.TrimSpace(description)
	d = strings.ReplaceAll(d, " ", "_")
	d = nonFileChar.ReplaceAllString(d, "")
	if d == "" {
		d = "memo"
	}
	return fmt.Sprintf("%s_%s.md", time.Now().Format("20060102"), d)
}
