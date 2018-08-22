package firebase

import (
	"github.com/labstack/echo"
	"strings"
	"os"
	"context"
	"firebase.google.com/go"
	"google.golang.org/api/option"
	"net/http"
	"firebase.google.com/go/auth"
)

var (
	opt = option.WithCredentialsFile(os.Getenv("CREDENTIALS"))
	authClient *auth.Client
)

// Firebaseの初期化処理
func Init() error {
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	authClient, err = app.Auth(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// Echo用Middleware
func JWTHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if _, err := GetUserToken(authHeader); err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "error verifying ID token"})
		}
		return next(c)
	}
}

// AuthorizationヘッダーからユーザーのToken情報を取得する
func GetUserToken(authHeader string) (*auth.Token, error) {
	idToken := strings.Replace(authHeader, "Bearer ", "", 1)
	return authClient.VerifyIDToken(context.Background(), idToken)
}

func GetUserName(uid string) (string, error) {
	record, err := authClient.GetUser(context.Background(), uid)
	if err != nil {
		return "", err
	}

	return record.DisplayName, nil
}