package request

type User struct {
	ID        string `json:"-"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FullName  string `json:"full_name"`
	CreatedAt string `json:"-"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
