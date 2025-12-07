package service

import (
	"strings"
)

type HeaderAnalyzer struct{}

func NewHeaderAnalyzer() *HeaderAnalyzer {
	return &HeaderAnalyzer{}
}

func (ha *HeaderAnalyzer) Analyze(text string) string {
	if text == "" {
		return ""
	}

	lines := strings.Split(text, "\n")

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		if trimmed == "" {
			continue
		}

		content := line

		if i+1 < len(lines) && strings.TrimSpace(lines[i+1]) != "" {
			if dotIndex := strings.Index(content, "."); dotIndex != -1 {
				content = content[:dotIndex+1]
			}
		}

		lines[i] = "<h1>" + content + "</h1>"
		break
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		trimmed := strings.TrimSpace(line)

		if trimmed == "" {
			continue
		}

		if !strings.HasPrefix(trimmed, "<strong>") || !strings.HasSuffix(trimmed, "</strong>") {
			continue
		}

		prevEmpty := i == 0 || strings.TrimSpace(lines[i-1]) == ""
		nextEmptyOrNotBold := i+1 >= len(lines) ||
			strings.TrimSpace(lines[i+1]) == "" ||
			!strings.HasPrefix(strings.TrimSpace(lines[i+1]), "<strong>")

		if prevEmpty && nextEmptyOrNotBold {
			lines[i] = "<h2>" + trimmed + "</h2>"
		}
	}

	return strings.Join(lines, "\n")
}
