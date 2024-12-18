package dto

type SignInInput struct {
	Email    string
	Password string
}

type SignInRes struct {
	Token string
}
