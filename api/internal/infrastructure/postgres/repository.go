package postgres

import (
	"app/internal/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) CreateOrUpdate(ctx context.Context, post *domain.Post) error {
	query := `
		INSERT INTO posts (id, group_id, title, url, text, seo_title, seo_description, seo_keywords, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())
		ON CONFLICT (id) DO UPDATE SET
			group_id = EXCLUDED.group_id,
			title = EXCLUDED.title,
			url = EXCLUDED.url,
			text = EXCLUDED.text,
			seo_title = EXCLUDED.seo_title,
			seo_description = EXCLUDED.seo_description,
			seo_keywords = EXCLUDED.seo_keywords,
			updated_at = NOW()
	`
	_, err := r.db.Exec(ctx, query,
		post.ID,
		post.GroupID,
		post.Title,
		post.URL,
		post.Text,
		post.SEOTitle,
		post.SEODescription,
		post.SEOKeywords,
		post.CreatedAt,
	)
	
	return err
}

func (r *PostRepository) GetList(ctx context.Context, page, limit int) ([]domain.Post, int, error) {
	offset := (page - 1) * limit
	countQuery := `SELECT COUNT(*) FROM posts`

	var count int

	err := r.db.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, group_id, title, url, text, seo_title, seo_description, seo_keywords, created_at, updated_at
		FROM posts
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var posts []domain.Post

	for rows.Next() {
		var post domain.Post

		err := rows.Scan(
			&post.ID,
			&post.GroupID,
			&post.Title,
			&post.URL,
			&post.Text,
			&post.SEOTitle,
			&post.SEODescription,
			&post.SEOKeywords,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}

		posts = append(posts, post)
	}

	return posts, count, nil
}
