package xecho

type HandlerFunc func(c XContext) error
type MiddlewareFunc func(next HandlerFunc) HandlerFunc
