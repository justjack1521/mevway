package database

import (
	"context"
	"fmt"
	"mevway/internal/adapter/database/dto"
	"mevway/internal/core/domain/content"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func (r *ArticleRepository) QueryNewsArticle(ctx context.Context, id uuid.UUID) (content.NewsArticle, error) {
	//TODO implement me
	panic("implement me")
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) QueryNewsArticleContainers(ctx context.Context, articleID uuid.UUID) ([]content.NewsContainer, error) {
	var rows []dto.ArticleNodeContainer
	err := r.db.WithContext(ctx).
		Where("article_id = ?", articleID).
		Order("sort_order").
		Find(&rows).Error
	if err != nil {
		return nil, fmt.Errorf("fetching containers: %w", err)
	}

	containers := make([]content.NewsContainer, len(rows))
	for i, row := range rows {
		containers[i] = row.ToEntity()
	}
	return containers, nil
}

func (r *ArticleRepository) QueryNewsArticleNodes(ctx context.Context, containerID uuid.UUID) ([]content.NewsNode, error) {
	var rows []dto.ArticleNodeGorm
	err := r.db.WithContext(ctx).
		Where("container_id = ?", containerID).
		Order("sort_order").
		Find(&rows).Error
	if err != nil {
		return nil, fmt.Errorf("fetching nodes: %w", err)
	}

	nodes := make([]content.NewsNode, 0, len(rows))
	for _, row := range rows {
		node, err := row.ToEntity()
		if err != nil {
			return nil, fmt.Errorf("converting node %s: %w", row.SysID, err)
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}
