package resources

type BanUserRequest struct {
	SysUser string `json:"SysUser" form:"SysUser" binding:"required"`
}
