package resources

type UserLoginRequest struct {
	Username   string `json:"Username" form:"Username" binding:"required"`
	Password   string `json:"Password" form:"Password" binding:"required"`
	RememberMe bool   `json:"RememberMe" form:"RememberMe"`
}

type UserLoginResponse struct {
	SysUser      string `json:"SysUser" form:"SysUser" binding:"required"`
	CustomerID   string `json:"CustomerID" form:"CustomerID" binding:"required"`
	AccessToken  string `json:"AccessToken" form:"AccessToken" binding:"required"`
	RefreshToken string `json:"RefreshToken" form:"RefreshToken" binding:"required"`
}
