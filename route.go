package xecho

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type xRoute struct {
	*echo.Route
	xe *XEcho
	h  echo.HandlerFunc
	m  []echo.MiddlewareFunc
}

func (xe *XEcho) CONNECT(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *xRoute {
	_h := func(c echo.Context) error {
		return h(*Context(xe, c))
	}

	return &xRoute{Route: xe.Add(http.MethodConnect, path, _h, m...), xe: xe, h: _h, m: m}
}

func (xe *XEcho) DELETE(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *xRoute {
	_h := func(c echo.Context) error {
		return h(*Context(xe, c))
	}

	return &xRoute{Route: xe.Add(http.MethodDelete, path, _h, m...), xe: xe, h: _h, m: m}
}

func (xe *XEcho) GET(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *xRoute {
	_h := func(c echo.Context) error {
		return h(*Context(xe, c))
	}

	return &xRoute{Route: xe.Add(http.MethodGet, path, _h, m...), xe: xe, h: _h, m: m}
}

func (xe *XEcho) HEAD(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *xRoute {
	_h := func(c echo.Context) error {
		return h(*Context(xe, c))
	}

	return &xRoute{Route: xe.Add(http.MethodHead, path, _h, m...), xe: xe, h: _h, m: m}
}

func (xe *XEcho) OPTIONS(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *xRoute {
	_h := func(c echo.Context) error {
		return h(*Context(xe, c))
	}

	return &xRoute{Route: xe.Add(http.MethodOptions, path, _h, m...), xe: xe, h: _h, m: m}
}

func (xe *XEcho) PATCH(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *xRoute {
	_h := func(c echo.Context) error {
		return h(*Context(xe, c))
	}

	return &xRoute{Route: xe.Add(http.MethodPatch, path, _h, m...), xe: xe, h: _h, m: m}
}

func (xe *XEcho) POST(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *xRoute {
	_h := func(c echo.Context) error {
		return h(*Context(xe, c))
	}

	return &xRoute{Route: xe.Add(http.MethodPost, path, _h, m...), xe: xe, h: _h, m: m}
}

func (xe *XEcho) PUT(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *xRoute {
	_h := func(c echo.Context) error {
		return h(*Context(xe, c))
	}

	return &xRoute{Route: xe.Add(http.MethodPut, path, _h, m...), xe: xe, h: _h, m: m}
}

func (xe *XEcho) TRACE(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *xRoute {
	_h := func(c echo.Context) error {
		return h(*Context(xe, c))
	}

	return &xRoute{Route: xe.Add(http.MethodTrace, path, _h, m...), xe: xe, h: _h, m: m}
}

func (xr *xRoute) Authenticated() *xRoute {
	if xr.xe.a == nil {
		log.Printf("Route %s cannot be authenticated since the authorization middleware is nil. Consider using WithAuthorize function for XEcho.", xr.Path)
		return xr
	}

	xr.Path = strings.TrimPrefix(xr.Path, "/")
	xr.Route = xr.xe.a.Add(xr.Method, xr.Path, xr.h, xr.m...)
	return xr
}

func (xr *xRoute) WithScopes(scope ...Stringer) *xRoute {
	if xr.xe.a == nil {
		log.Printf("Route %s cannot have scopes since the authorization middleware is nil. Consider using WithAuthorize function for XEcho.", xr.Path)
		return xr
	}

	xr.m = append(xr.m, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			xc := Context(xr.xe, c)
			ok := scopeCheck(xc, scope...)
			if !ok {
				return xc.XError(
					NewError(*xc).WithResponseCode(http.StatusUnauthorized).
						WithCode("no_permission"),
				)
			}

			return next(c)
		}
	})

	xr.Path = strings.TrimPrefix(xr.Path, "/")
	xr.Route = xr.xe.a.Add(xr.Method, xr.Path, xr.h, xr.m...)
	return xr
}

func scopeCheck(c *XContext, epScope ...Stringer) bool {
	var (
		tokenScope, ok = c.Scope()
		tokenScopeMap  = make(map[string]any)
	)

	if len(epScope) == 0 {
		return false
	}

	if !ok {
		return false
	}

	// If the token has a wildcard scope, then it has access to all scopes.
	// This is useful for authenticated users.
	if tokenScope[0] == "*" {
		return true
	}

	for _, s := range tokenScope {
		tokenScopeMap[s] = struct{}{}
	}

	for _, s := range epScope {
		if _, ok := tokenScopeMap[s.String()]; !ok {
			return false
		}
	}

	return true
}
