package posts

import (
	"context"
	"fmt"

	"app/internal/domain"
)

type GetListUseCase struct {
	repository domain.PostRepository
}

type PostSnippet struct {
	ID        int64   `json:"id"`
	GroupID   int64   `json:"group_id"`
	Title     *string `json:"title,omitempty"`
	URL       *string `json:"url,omitempty"`
	Text      *string `json:"text,omitempty"` // snippet
	CreatedAt string  `json:"created_at"`
}

type ListResponse struct {
	Items    []PostSnippet `json:"items"`
	Paginate struct {
		Page  int `json:"page"`
		Count int `json:"count"`
		Limit int `json:"limit"`
	} `json:"paginate"`
}

func NewListPostsUseCase(postRepository domain.PostRepository) *GetListUseCase {
	return &GetListUseCase{
		repository: postRepository,
	}
}

func (uc *GetListUseCase) Execute(ctx context.Context, page int, limit int) (*ListResponse, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 9
	} else if limit > 100 {
		limit = 100
	}

	posts, count, err := uc.repository.GetList(ctx, page, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to list posts: %w", err)
	}

	items := make([]PostSnippet, len(posts))

	for i, post := range posts {
		snippet := *post.Text
		runes := []rune(snippet)

		if len(runes) > 200 {
			snippet = string(runes[:200])
		}

		items[i] = PostSnippet{
			ID:        post.ID,
			GroupID:   post.GroupID,
			Title:     post.Title,
			URL:       post.URL,
			Text:      &snippet,
			CreatedAt: post.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
	}

	response := &ListResponse{
		Items: items,
	}

	response.Paginate.Page = page
	response.Paginate.Count = count
	response.Paginate.Limit = limit

	return response, nil
}
