package service

import (
	"regexp"
	"strings"
)

type TitleExtractor struct{}

func NewTitleExtractor() *TitleExtractor {
	return &TitleExtractor{}
}

func (te *TitleExtractor) Extract(text string) (string, string) {
	if text == "" {
		return "", ""
	}

	h1Regex := regexp.MustCompile(`(?is)<h1[^>]*>(.*?)</h1>`)
	match := h1Regex.FindStringSubmatch(text)

	if len(match) < 2 {
		return "", text
	}

	htmlRegex := regexp.MustCompile(`<[^>]*>`)
	title := strings.TrimSpace(match[1])
	title = htmlRegex.ReplaceAllString(title, "")
	title = strings.TrimSpace(title)

	modifiedText := h1Regex.ReplaceAllString(text, "")
	modifiedText = strings.TrimSpace(modifiedText)

	return title, modifiedText
}
