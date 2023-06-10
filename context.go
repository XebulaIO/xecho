package xecho

import (
	"github.com/labstack/echo/v4"
)

var (
	sessionKey = "session-key"
)

type xContext[T any] struct {
	echo.Context
}

func Context[T any](c echo.Context) *xContext[T] {
	return &xContext[T]{c}
}

func ContextWithSession[T any](c echo.Context, s *T) *xContext[T] {
	return &xContext[T]{c}
}

func (c *xContext[T]) Session() *T {
	s := c.Get(sessionKey)
	if s == nil {
		return nil
	}

	return s.(*T)
}

func (c *xContext[T]) SetSession(s *T) {
	c.Set(sessionKey, s)
}
