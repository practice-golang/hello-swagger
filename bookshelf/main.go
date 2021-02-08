package main // import "bookshelf"

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"auth"
	"books"
	"commons"
	"consts"

	// echoSwagger "github.com/swaggo/echo-swagger"
	echoSwagger "echo-swagger"

	_ "docs"
)

var (
	isTest = false
	file   *os.File
)

func dumpHandler(c echo.Context, reqBody, resBody []byte) {
	header := time.Now().Format("2006-01-02 15:04:05") + " - "
	body := string(reqBody)
	body = strings.Replace(body, "\r\n", "", -1)
	body = strings.Replace(body, "\n", "", -1)
	data := header + body + "\n"

	f, err := os.OpenFile(
		"request-body.log",
		os.O_APPEND|os.O_CREATE|os.O_RDWR,
		os.FileMode(0777),
	)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	if _, err = f.WriteString(data); err != nil {
		log.Println(err)
		return
	}
}

func serverSetup() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Use(
		middleware.CORS(),
		middleware.Recover(),
	)

	d := e.Group("")
	d.Use(middleware.BasicAuth(auth.DocsLogin))
	d.GET("/docs/*", echoSwagger.WrapHandler)
	url := echoSwagger.URL("./doc.json")
	d.GET("/docs/*", echoSwagger.EchoWrapHandler(url))

	// Route - Public
	p := e.Group("/api/v1")
	p.GET("/health", commons.Health)
	p.GET("/sample-token", commons.GenerateSampleToken)

	// Route - Auth
	config := middleware.JWTConfig{
		Claims:        &commons.CustomClaims{},
		SigningKey:    consts.JWTsecret,
		SigningMethod: "HS256",
	}

	g := e.Group("/api/v1")
	g.Use(middleware.JWTWithConfig(config))
	g.GET("/get-book/:page", books.GetBook)
	g.POST("/search-book", books.SearchBook)
	g.POST("/add-book", books.AddBook)

	return e
}

// @title Bookshelf API
// @version 0.0.1
// @description My sample bookshelf server.
// @description
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	e := serverSetup()

	if !isTest {
		var err error
		file, err = os.OpenFile(
			"connection.log",
			os.O_APPEND|os.O_CREATE|os.O_RDWR,
			os.FileMode(0777),
		)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: `${time_rfc3339} - remote_ip:${remote_ip}, host:${host}, ` +
				`method:${method}, uri:${uri},status:${status}, error:${error}, ` +
				`${header:Authorization}, query:${query:property}, form:${form}, ` + "\n",
			Output: file,
		}))

		e.Use(middleware.BodyDump(dumpHandler))
	}

	e.Logger.Fatal(e.Start(":4416"))
}
