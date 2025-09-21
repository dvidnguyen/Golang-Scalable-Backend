package usecase

type EmailPasswordRegistration struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}
type EmailPasswordLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
