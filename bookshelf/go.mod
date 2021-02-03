module bookshelf

go 1.15

require (
	auth v0.0.0
	books v0.0.0
	commons v0.0.0
	consts v0.0.0
	docs v0.0.0
	github.com/appleboy/gofight/v2 v2.1.2
	github.com/labstack/echo/v4 v4.1.17
	github.com/stretchr/testify v1.7.0
	github.com/swaggo/echo-swagger v1.1.0
)

replace (
	auth => ./auth
	books => ./books
	commons => ./commons
	consts => ./consts
	docs => ./docs
)
