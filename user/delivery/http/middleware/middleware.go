package middleware

import (
	"log"
	"macaiki/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

func JWTCreateToken(userId int, role string) (string, error) {
	JWTSecret, err := config.LoadJWTSecret(".")
	if err != nil {
		log.Fatal("err", err)
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTSecret.Secret))
}

func ExtractTokenUser(c echo.Context) (int, string) {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		role := claims["role"].(string)
		return int(userId), role
	}
	return 0, ""
}

// InitMiddleware initialize the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
