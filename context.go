package xecho

import (
	"github.com/labstack/echo/v4"
)

var (
	sessionKey = "session-key"
)

type XContext struct {
	echo.Context
}

func Context(c echo.Context) *XContext {
	return &XContext{c}
}

func (c *XContext) Session() Session {
	s := c.Get(sessionKey)
	if s == nil {
		return nil
	}

	return s.(Session)
}

func (c *XContext) SetSession(s Session) {
	c.Set(sessionKey, s)
}
