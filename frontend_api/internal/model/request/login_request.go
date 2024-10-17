package request

type LoginRequest struct {
	Email    string `json:"email,omitempty"`
	Account  string `json:"account,omitempty"`
	Password string `json:"password"`
}
