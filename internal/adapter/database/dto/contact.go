package dto

type ContactGorm struct {
	Email   string `gorm:"column:email"`
	Content string `gorm:"column:content"`
}

func (ContactGorm) TableName() string {
	return "system.contact"
}
