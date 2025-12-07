package service

import (
	"regexp"
	"strings"

	"github.com/essentialkaos/translit"
)

type URLNormalizer struct {
}

func NewURLNormalizer() *URLNormalizer {
	return &URLNormalizer{}
}

func (s *URLNormalizer) Normalize(title string) string {
	if title == "" {
		return ""
	}

	transliterated := translit.EncodeToICAO(title)
	cleaned := strings.ToLower(transliterated)
	cleaned = regexp.MustCompile(`[\s\-_]+`).ReplaceAllString(cleaned, "-")
	cleaned = regexp.MustCompile(`[^a-z0-9\-_]`).ReplaceAllString(cleaned, "")
	cleaned = strings.Trim(cleaned, "-")
	cleaned = regexp.MustCompile(`-+`).ReplaceAllString(cleaned, "-")

	return cleaned
}
