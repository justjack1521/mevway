package dto

type IPAddressBlacklistGorm struct {
	ID      uint   `gorm:"column:id;primaryKey;autoIncrement"`
	IPRange string `gorm:"column:ip_range;type:inet"`
}

func (IPAddressBlacklistGorm) TableName() string {
	return "system.ip_address_blacklist"
}
