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

func (xe *XEcho) UseAuthenticated(f MiddlewareFunc) *XEcho {
	if xe.a == nil {
		xe.a = xe.Group("/")
	}

	mw := func(next echo.HandlerFunc) echo.HandlerFunc {
		r := f(func(xc XContext) error {
			return next(xc)
		})

		return func(c echo.Context) error {
			return r(*Context(xe, c))
		}
	}

	xe.a.Use(mw)
	return xe
}
