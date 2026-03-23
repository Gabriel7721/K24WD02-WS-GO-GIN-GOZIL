package auth

type LoginInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
