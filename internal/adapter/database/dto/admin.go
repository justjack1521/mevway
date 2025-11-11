package dto

type IPAddressBlacklistGorm struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	IPRange string `gorm:"type:inet"`
}

func (IPAddressBlacklistGorm) TableName() string {
	return "system.ip_address_blacklist"
}
