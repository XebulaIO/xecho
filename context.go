package xecho

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var (
	sessionKey = "session-key"
	scopeKey   = "scope"
)

type XContext struct {
	echo.Context

	e  *XEcho
	id string
}

func Context(e *XEcho, c echo.Context) *XContext {
	return &XContext{Context: c, e: e, id: uuid.New().String()}
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

func (c *XContext) ID() string {
	return c.id
}

func (c *XContext) XError(err *Error) error {
	logger.Error(err)
	return c.JSON(err.responseCode, err)
}

func (c *XContext) XEcho() *XEcho {
	return c.e
}

func (c *XContext) SetScope(scope []string) {
	c.Set(scopeKey, scope)
}

func (c *XContext) Scope() ([]string, bool) {
	scope := c.Get(scopeKey)
	if scope == nil {
		return nil, false
	}

	s, ok := scope.([]string)
	return s, ok
}
