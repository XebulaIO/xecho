package xecho

import (
	"github.com/labstack/echo/v4"
)

type XEcho struct {
	*echo.Echo
	a *echo.Group
}

func New(e *echo.Echo) *XEcho {
	return &XEcho{
		Echo: e,
	}
}

func (xe *XEcho) WithAuthorize(authorize echo.MiddlewareFunc) *XEcho {
	xe.a = xe.Group("/", authorize)
	return xe
}
