package xecho

type Session interface {
	Delete() error
	Token() string
}
