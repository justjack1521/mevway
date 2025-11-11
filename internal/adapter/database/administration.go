package database

import (
	"context"
	"gorm.io/gorm"
	"mevway/internal/adapter/database/dto"
	"net"
)

type AdministrationRepository struct {
	database *gorm.DB
}

func NewAdministrationRepository(database *gorm.DB) *AdministrationRepository {
	return &AdministrationRepository{database: database}
}

func (r *AdministrationRepository) IPAddressBlacklisted(ctx context.Context, ip net.IP) (bool, error) {
	var count int64
	if err := r.database.WithContext(ctx).Model(&dto.IPAddressBlacklistGorm{}).Where("ip_range >>= ?", ip.String()).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
