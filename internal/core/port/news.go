package port

import (
	"context"
	"mevway/internal/core/domain/content"

	uuid "github.com/satori/go.uuid"
)

type NewsRepository interface {
	QueryNewsArticle(ctx context.Context, id uuid.UUID) (content.NewsArticle, error)
}

type NewsService interface {
	GetNewsArticle(ctx context.Context, id uuid.UUID) (content.NewsArticle, error)
}
