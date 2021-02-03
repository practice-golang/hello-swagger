module commons

go 1.15

require (
	consts v0.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/labstack/echo/v4 v4.1.17
	github.com/stretchr/testify v1.7.0
)

replace (
	consts => ../consts
)
