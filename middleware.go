package xecho

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type HandlerFunc func(c XContext) error
type MiddlewareFunc func(next HandlerFunc) HandlerFunc
type ValidateTokenFunc func(string) (Session, error)

type BearerAuthenticationMiddleWareConfig struct {
	CookieName    string
	ValidateToken ValidateTokenFunc
}

type JWTCasbinAuthorizationMiddlewareConfig struct {
	Enforcer *casbin.Enforcer
	GetRoles func(sub string) ([]string, error)
	KeyFunc  func(token *jwt.Token) (interface{}, error)
}

func BearerAuthenticationMiddleWareWithConfig(config BearerAuthenticationMiddleWareConfig) MiddlewareFunc {
	if config.ValidateToken == nil {
		config.ValidateToken = func(s string) (Session, error) { return nil, nil }
	}

	getToken := func(c echo.Context) string {
		authToken := c.Request().Header.Get(echo.HeaderAuthorization)
		if authToken == "" {
			cookie, err := c.Cookie(config.CookieName)
			if err != nil {
				return ""
			}

			return cookie.Value
		}

		splitToken := strings.Split(authToken, "Bearer ")
		if len(splitToken) > 1 {
			return splitToken[1]
		}

		return ""
	}

	return func(next HandlerFunc) HandlerFunc {
		return func(c XContext) error {
			token := getToken(c)
			if token == "" {
				c.NoContent(http.StatusUnauthorized)
				return fmt.Errorf("token nil")
			}

			s, err := config.ValidateToken(token)
			if err != nil {
				c.NoContent(http.StatusUnauthorized)
				return err
			}

			c.SetSession(s)

			if err := next(c); err != nil {
				c.Error(err)
			}

			return nil
		}
	}
}

func JWTCasbinAuthorizationMiddlewareWithConfig(config JWTCasbinAuthorizationMiddlewareConfig) MiddlewareFunc {
	if config.KeyFunc == nil {
		config.KeyFunc = func(token *jwt.Token) (interface{}, error) {
			return "", nil
		}
	}

	if config.GetRoles == nil {
		config.GetRoles = func(sub string) ([]string, error) {
			return []string{}, nil
		}
	}

	return func(next HandlerFunc) HandlerFunc {
		return func(c XContext) error {
			var (
				allow bool
				s     = c.Session()
			)

			claims := jwt.MapClaims{}
			_, err := jwt.ParseWithClaims(s.AccessToken(), claims, config.KeyFunc)
			if err != nil {
				return err
			}

			sub, ok := claims["sub"].(string)
			if !ok {
				return fmt.Errorf("invalid token: sub not found")
			}

			roles, err := config.GetRoles(sub)
			if err != nil {
				return err
			}

			for _, r := range roles {
				allow, _ = config.Enforcer.Enforce(r, c.Path(), strings.ToLower(c.Request().Method))
				if allow {
					break
				}
			}

			if !allow {
				c.NoContent(http.StatusUnauthorized)
				return fmt.Errorf("no_permission")
			}

			return next(c)
		}
	}
}
