module books

go 1.15

require (
	github.com/labstack/echo/v4 v4.1.17
	github.com/stretchr/testify v1.7.0
	commons v0.0.0
)

replace (
	commons => ../commons
	consts => ../consts
)
