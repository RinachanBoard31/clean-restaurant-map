package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func SetupJwtMiddleware(g *echo.Group) {
	g.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("JWT_SIGNING_KEY")),
		TokenLookup: "cookie:" + os.Getenv("JWT_TOKEN_NAME"),
		SuccessHandler: func(c echo.Context) {
			// 本来はここでJWTのclaimsが取得できるができない時があったため、取得できない時は手動で認証を確認して問題なければclaimsを取得する
			rowClaims := c.Get("claims")
			var claims jwt.MapClaims = nil
			if rowClaims != nil {
				claims = rowClaims.(jwt.MapClaims)
			}
			if rowClaims == nil { // 手動で認証する
				var err error = nil
				claims, err = validateJwtOnSelf(c.Cookie(os.Getenv("JWT_TOKEN_NAME")))
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			userId := claims["sub"].(string)
			c.Set("userId", userId)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			// 認証が失敗したらリダイレクト
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Unauthorized",
			})
		},
	}))
}

// echojwt.WithConfigでJETが合っているはずなのにtoken.Validでエラーが出る時があるため手動で確認する
func validateJwtOnSelf(cookie *http.Cookie, err error) (jwt.MapClaims, error) {
	if err != nil {
		return nil, err
	}
	tokenString := cookie.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("unexpected signing method")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("JWT_SIGNING_KEY"), nil
	})
	if err != nil {
		fmt.Println("Error parsing token:", err)
	}
	claims := token.Claims.(jwt.MapClaims)
	if claims["exp"] != nil {
		expTime := int64(claims["exp"].(float64))
		if expTime < time.Now().Unix() {
			fmt.Println("Token expired")
		}
	}
	jwtData := strings.Split(tokenString, ".")
	if len(jwtData) != 3 {
		return nil, fmt.Errorf("jwtData is not included header, payload or signature")
	}
	// ヘッダーとペイロードを結合
	encodedData := jwtData[0] + "." + jwtData[1]
	secret := []byte(os.Getenv("JWT_SIGNING_KEY"))
	signature := createSignature(encodedData, secret)
	if jwtData[2] == signature {
		return claims, nil
	}
	return nil, fmt.Errorf("authentication failure")
}

// Base64エンコード
func base64UrlEncode(input []byte) string {
	encoded := base64.StdEncoding.EncodeToString(input)
	// Base64Urlエンコード（+ → - , / → _ , = を削除）
	encoded = strings.TrimRight(encoded, "=")
	encoded = strings.ReplaceAll(encoded, "+", "-")
	encoded = strings.ReplaceAll(encoded, "/", "_")
	return encoded
}

// HMAC-SHA256で署名を作成
func createSignature(data string, secret []byte) string {
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(data))
	return base64UrlEncode(mac.Sum(nil))
}
