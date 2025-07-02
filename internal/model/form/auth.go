package form

type RegisterForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
