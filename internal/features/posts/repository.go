package posts

import (
	"context"
	"database/sql"
	"fmt"
)

const postsPerPage = 10

type Repository interface {
	List(ctx context.Context, page int) ([]Post, int, error)
}

type PostRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) List(ctx context.Context, page int) ([]Post, int, error) {
	var total int
	if err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM posts`).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("contar posts: %w", err)
	}

	offset := (page - 1) * postsPerPage
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, fk_user_id, title, message
		FROM posts
		ORDER BY id
		LIMIT $1 OFFSET $2`, postsPerPage, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("consultar posts: %w", err)
	}
	defer rows.Close()

	posts := make([]Post, 0, postsPerPage)
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.FKUserID, &post.Title, &post.Message); err != nil {
			return nil, 0, fmt.Errorf("ler post: %w", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("iterar posts: %w", err)
	}

	return posts, total, nil
}
