package xecho

type Session interface {
	Delete() error
	AccessToken() string
	RefreshToken() string
}
