package auth

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int
}
