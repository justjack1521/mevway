package application

import (
	"context"
	"mevway/internal/core/domain/content"
	"mevway/internal/core/port"

	uuid "github.com/satori/go.uuid"
)

type NewsArticleService struct {
	repository port.NewsRepository
}

func NewNewsArticleService(repository port.NewsRepository) *NewsArticleService {
	return &NewsArticleService{repository: repository}
}

func (s *NewsArticleService) GetNewsArticle(ctx context.Context, id uuid.UUID) (content.NewsArticle, error) {

	article, err := s.repository.QueryNewsArticle(ctx, id)
	if err != nil {
		return content.NewsArticle{}, err
	}

	containers, err := s.repository.QueryNewsArticleContainers(ctx, article.ID)
	if err != nil {
		return content.NewsArticle{}, err
	}
	article.Containers = containers

	for index, container := range article.Containers {
		nodes, err := s.repository.QueryNewsArticleNodes(ctx, container.ID)
		if err != nil {
			return content.NewsArticle{}, err
		}
		article.Containers[index].Nodes = nodes
	}

	return article, nil

}
