package sync

import (
	"context"
	"fmt"
	"strconv"

	"app/internal/domain"
	"app/internal/domain/service"
)

type SyncRawMessageUseCase struct {
	postRepository domain.PostRepository
	textNormalizer *service.TextNormalizer
	headerAnalyzer *service.HeaderAnalyzer
	titleExtractor *service.TitleExtractor
	urlNormalizer  *service.URLNormalizer
}

func NewSyncRawMessageUseCase(postRepository domain.PostRepository) *SyncRawMessageUseCase {
	return &SyncRawMessageUseCase{
		postRepository: postRepository,
		textNormalizer: service.NewTextNormalizer(),
		headerAnalyzer: service.NewHeaderAnalyzer(),
		titleExtractor: service.NewTitleExtractor(),
		urlNormalizer:  service.NewURLNormalizer(),
	}
}

func (uc *SyncRawMessageUseCase) Execute(ctx context.Context, rawMessage domain.RawMessage) error {
	var normalizedText *string
	var title *string
	var url *string

	if rawMessage.Text != nil {
		normalized := uc.textNormalizer.Normalize(*rawMessage.Text, rawMessage.Entities)
		analyzed := uc.headerAnalyzer.Analyze(normalized)

		extractedTitle, modifiedText := uc.titleExtractor.Extract(analyzed)

		if extractedTitle != "" {
			title = &extractedTitle
			cleanedURL := uc.urlNormalizer.Normalize(extractedTitle)

			if cleanedURL != "" {
				combinedURL := strconv.Itoa(rawMessage.ID) + "-" + cleanedURL
				url = &combinedURL
			}
		}

		normalizedText = &modifiedText
	}

	post := &domain.Post{
		ID:        int64(rawMessage.ID),
		GroupID:   rawMessage.GroupID,
		Title:     title,
		URL:       url,
		Text:      normalizedText,
		CreatedAt: rawMessage.Date,
		// SEOTitle, SEODescription, SEOKeywords can be set later or generated
	}

	if post.URL == nil && post.GroupID == 0 {
		return nil
	}

	err := uc.postRepository.CreateOrUpdate(ctx, post)
	if err != nil {
		return fmt.Errorf("failed to save post from raw message %d: %w", rawMessage.ID, err)
	}

	return nil
}
