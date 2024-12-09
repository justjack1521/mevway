package database

import (
	"context"
	"gorm.io/gorm"
	"mevway/internal/adapter/database/dto"
)

type ContactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) Create(ctx context.Context, email string, content string) error {
	var result = &dto.ContactGorm{
		Email:   email,
		Content: content,
	}
	if err := r.db.WithContext(ctx).Model(result).Create(result).Error; err != nil {
		return err
	}
	return nil
}
