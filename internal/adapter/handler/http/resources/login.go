package resources

type UserLoginRequest struct {
	Username   string `json:"Username" form:"Username" binding:"required"`
	Password   string `json:"Password" form:"Password" binding:"required"`
	RememberMe bool   `json:"RememberMe" form:"RememberMe"`
}

type UserLoginResponse struct {
	SessionID     string `json:"SessionID" form:"SessionID"`
	IDToken       string `json:"IDToken" form:"IDToken" binding:"required"`
	AccessToken   string `json:"AccessToken" form:"AccessToken" binding:"required"`
	RefreshToken  string `json:"RefreshToken" form:"RefreshToken" binding:"required"`
	RememberToken string `json:"RememberToken" form:"RememberToken"`
}
