package commons // import "commons"

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"consts"
)

// CustomClaims : Custom claims
type CustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

// Health - 서버 동작여부
// @Tags Etc.
// @Summary Server status
// @Produce json
// @Success 200 {object} string "OK"
// @Router /health [get]
func Health(c echo.Context) error {
	result := "OK"
	return c.String(http.StatusOK, result)
}

// GenerateSampleToken - 샘플 토큰
// @Tags Etc.
// @Summary Get sample token for test
// @Produce json
// @Success 200 {string} string "{"token": "..."}"
// @Router /sample-token [get]
func GenerateSampleToken(c echo.Context) error {
	// claims := &models.CustomClaims{"Sample User", jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 7200).Unix()}}
	claims := &CustomClaims{"Sample User", jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 7200).Unix()}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(consts.JWTsecret)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusOK, echo.Map{"token": "Bearer " + t})
}
