package resources

type RememberUserRequest struct {
	Token string `json:"Token" form:"Token" binding:"required"`
}

type RememberUserResponse struct {
	SessionID    string `json:"SessionID" form:"SessionID" binding:"required"`
	CustomerID   string `json:"CustomerID" form:"CustomerID" binding:"required"`
	Username     string `json:"Username" form:"Username" binding:"required"`
	AccessToken  string `json:"AccessToken" form:"AccessToken" binding:"required"`
	RefreshToken string `json:"RefreshToken" form:"RefreshToken" binding:"required"`
}
