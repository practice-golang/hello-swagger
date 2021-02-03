package auth // import "auth"

import "github.com/labstack/echo/v4"

// DocsLogin : Swagger 진입시 BasicAuth
func DocsLogin(username, password string, c echo.Context) (bool, error) {
	if username == "dev" && password == "1234" {
		return true, nil
	}
	return false, nil
}
