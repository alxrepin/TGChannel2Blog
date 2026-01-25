package text

import (
	"regexp"
	"strings"

	"github.com/essentialkaos/translit"
)

var (
	reSpaces  = regexp.MustCompile(`[\s\-_]+`)
	reInvalid = regexp.MustCompile(`[^a-z0-9\-_]`)
	reDashes  = regexp.MustCompile(`-+`)
)

func Translit(text string) string {
	if text == "" {
		return ""
	}

	transliterated := translit.EncodeToICAO(text)
	cleaned := strings.ToLower(transliterated)
	cleaned = reSpaces.ReplaceAllString(cleaned, "-")
	cleaned = reInvalid.ReplaceAllString(cleaned, "")
	cleaned = strings.Trim(cleaned, "-")
	cleaned = reDashes.ReplaceAllString(cleaned, "-")

	return cleaned
}

func UTF16ToByteOffset(text string, utf16Pos int) int {
	runes := []rune(text)

	var bytePos int
	var currentUtf16 int

	for _, r := range runes {
		if currentUtf16 == utf16Pos {
			return bytePos
		}

		utf16Inc := 1

		if r > 0xFFFF {
			utf16Inc = 2
		}

		currentUtf16 += utf16Inc
		bytePos += len(string(r))
	}

	return bytePos
}
