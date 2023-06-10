package xecho

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type xRoute struct {
	*echo.Route
	xe *XEcho
	h  echo.HandlerFunc
	m  []echo.MiddlewareFunc
}

func (xe *XEcho) CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *xRoute {
	return &xRoute{Route: xe.Add(http.MethodConnect, path, h, m...), xe: xe, h: h, m: m}
}

func (xe *XEcho) DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *xRoute {
	return &xRoute{Route: xe.Add(http.MethodDelete, path, h, m...), xe: xe, h: h, m: m}
}

func (xe *XEcho) GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *xRoute {
	return &xRoute{Route: xe.Add(http.MethodGet, path, h, m...), xe: xe, h: h, m: m}
}

func (xe *XEcho) HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *xRoute {
	return &xRoute{Route: xe.Add(http.MethodHead, path, h, m...), xe: xe, h: h, m: m}
}

func (xe *XEcho) OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *xRoute {
	return &xRoute{Route: xe.Add(http.MethodOptions, path, h, m...), xe: xe, h: h, m: m}
}

func (xe *XEcho) PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *xRoute {
	return &xRoute{Route: xe.Add(http.MethodPatch, path, h, m...), xe: xe, h: h, m: m}
}

func (xe *XEcho) POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *xRoute {
	return &xRoute{Route: xe.Add(http.MethodPost, path, h, m...), xe: xe, h: h, m: m}
}

func (xe *XEcho) PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *xRoute {
	return &xRoute{Route: xe.Add(http.MethodPut, path, h, m...), xe: xe, h: h, m: m}
}

func (xe *XEcho) TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *xRoute {
	return &xRoute{Route: xe.Add(http.MethodTrace, path, h, m...), xe: xe, h: h, m: m}
}

func (xr *xRoute) Authenticated() *xRoute {
	if xr.xe.a == nil {
		log.Printf("Route %s cannot be authenticated since the authorization middleware is nil. Consider using WithAuthorize function for XEcho.", xr.Path)
		return xr
	}

	xr.Route = xr.xe.a.Add(xr.Method, xr.Path, xr.h, xr.m...)
	return xr
}
