package usecases

type CustomerSignUp interface {
	SignUp(name, email, password string) error
}
