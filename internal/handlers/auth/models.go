package auth

type UserCredentials struct {
	Username string
	Password string
}

type LoginParams struct {
	UserCredentials
}

type SignupParams struct {
	UserCredentials
	InitElo int16
}
