package auth

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	IDToken      string
	AccessToken  string
	RefreshToken string
	ExpiresIn    int
}
