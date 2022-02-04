package requests

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	UserID       int
	AccessToken  string
	RefreshToken string
}
