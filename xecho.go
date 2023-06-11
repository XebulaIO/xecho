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

func (xe *XEcho) WithAuthorize(authorize MiddlewareFunc) *XEcho {
	xe.a = xe.Group("/", func(next echo.HandlerFunc) echo.HandlerFunc {
		r := authorize(func(xc XContext) error {
			return next(xc)
		})

		return func(c echo.Context) error {
			return r(*Context(xe, c))
		}
	})

	return xe
}
