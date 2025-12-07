package service

import (
	"regexp"
	"sort"
	"strconv"
	"strings"

	"app/internal/domain"
)

type Tag struct {
	Opening string
	Closing string
}

type TagEvent struct {
	Pos       int
	IsOpening bool
	Tag       string
}

type TextNormalizer struct{}

func NewTextNormalizer() *TextNormalizer {
	return &TextNormalizer{}
}

func (tn *TextNormalizer) Normalize(text string, entities []domain.RawMessageEntity) string {
	if text == "" {
		return ""
	}
	if entities == nil || len(entities) == 0 {
		return text
	}

	sort.Slice(entities, func(i, j int) bool {
		return entities[i].Offset < entities[j].Offset
	})

	var events []TagEvent

	for _, entity := range entities {
		startByte := utf16ToByteOffset(text, entity.Offset)
		endByte := utf16ToByteOffset(text, entity.Offset+entity.Length)
		subtext := text[startByte:endByte]
		tag := getTag(entity.Type, entity, subtext)

		if tag.Opening != "" || tag.Closing != "" {
			events = append(events, TagEvent{Pos: startByte, IsOpening: true, Tag: tag.Opening})
			events = append(events, TagEvent{Pos: endByte, IsOpening: false, Tag: tag.Closing})
		}
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].Pos != events[j].Pos {
			return events[i].Pos < events[j].Pos
		}
		return events[i].IsOpening && !events[j].IsOpening
	})

	var builder strings.Builder
	pos := 0

	for _, event := range events {
		builder.WriteString(text[pos:event.Pos])
		builder.WriteString(event.Tag)
		pos = event.Pos
	}

	builder.WriteString(text[pos:])

	result := builder.String()
	re := regexp.MustCompile(`\n(</[^>]+>)`)
	result = re.ReplaceAllString(result, "$1\n")
	return result
}

func utf16ToByteOffset(text string, utf16Pos int) int {
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

func getTag(typ domain.RawMessageEntityType, entity domain.RawMessageEntity, text string) Tag {
	switch typ {
	case domain.EntityTypeBold:
		return Tag{"<strong>", "</strong>"}
	case domain.EntityTypeItalic:
		return Tag{"<em>", "</em>"}
	case domain.EntityTypeCode:
		return Tag{"<code>", "</code>"}
	case domain.EntityTypePre:
		return Tag{"<pre>", "</pre>"}
	case domain.EntityTypeTextLink:
		if entity.URL != nil {
			return Tag{`<a href="` + *entity.URL + `">`, "</a>"}
		}

		return Tag{"", ""}
	case domain.EntityTypeURL:
		if entity.URL != nil {
			return Tag{`<a href="` + *entity.URL + `">`, "</a>"}
		}

		return Tag{"", ""}
	case domain.EntityTypeCustomEmoji:
		if entity.CustomEmojiID != nil {
			return Tag{`<span data-emoji-id="` + strconv.FormatInt(*entity.CustomEmojiID, 10) + `">`, "</span>"}
		}

		return Tag{"", ""}
	default:
		return Tag{"", ""}
	}
}
