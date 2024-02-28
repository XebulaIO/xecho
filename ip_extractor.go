package xecho

import (
	"net"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func extractIP() echo.IPExtractor {
	return func(req *http.Request) string {
		realIP := req.Header.Get("Cf-Connecting-Ip")

		if realIP == "" {
			realIP = req.Header.Get(echo.HeaderXRealIP)
		}

		if realIP != "" {
			realIP = strings.TrimPrefix(realIP, "[")
			realIP = strings.TrimSuffix(realIP, "]")
			if ip := net.ParseIP(realIP); ip != nil {
				return realIP
			}
		}

		ra, _, _ := net.SplitHostPort(req.RemoteAddr)
		return ra
	}
}
