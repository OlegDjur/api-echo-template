package domain

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"type"`
	Password string `json:"password"`
}

type UserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"type"`
	Password string `json:"password"`
}
