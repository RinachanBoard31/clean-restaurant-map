package middleware

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func SetupJwtMiddleware(g *echo.Group) {
	g.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("JWT_SIGNING_KEY")),
		TokenLookup: "cookie:" + os.Getenv("JWT_TOKEN_NAME"),
		SuccessHandler: func(c echo.Context) {
			// JWTのClaimsのuserIdを取得し、コンテキストにセット
			claims := c.Get("claims").(*jwt.MapClaims)
			userId := (*claims)["sub"].(string)
			c.Set("userId", userId)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			// 認証が失敗したらリダイレクト
			url := os.Getenv("FRONT_URL")
			return c.Redirect(http.StatusFound, url)
		},
	}))
}
