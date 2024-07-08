package resources

import uuid "github.com/satori/go.uuid"

type UserRegisterRequest struct {
	Username        string `json:"Username" form:"Username" binding:"required"`
	Password        string `json:"Password" form:"Password" binding:"required"`
	ConfirmPassword string `json:"ConfirmPassword" form:"ConfirmPassword" binding:"required"`
}

type UserRegisterResponse struct {
	SysUser uuid.UUID `json:"SysUser" form:"SysUser" binding:"required"`
}
