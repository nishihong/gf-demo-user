package model

type UserCreateInput struct {
	Passport string
	Password string
	Nickname string
}

type UserSignInInput struct {
	Passport string
	Password string
}

type UserUpdateInput struct {
	Id       string
	Passport string
	Nickname string
}
