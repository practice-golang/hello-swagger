package books

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var mockBook = Book{
	Name:      "Justice: What's the Right Thing to Do?",
	Author:    "Michael Sandel",
	Edition:   2,
	Price:     9450,
	Unit:      "KRW",
	Publisher: "미래엔",
}
var mocBookJSON = `{"name":"기동전사 건담 섬광의 하사웨이","author":"토미노 요시유키","edition":1,"price":5000,"unit":"KRW","publisher":"시공사"}`

func TestGetBook(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/get-book/:page")
	c.SetParamNames("page")
	c.SetParamValues("1")

	if assert.NoError(t, GetBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, rec.Body.String(), `{"method":"POST","page":"1","where":"getbook"}`+"\n")
	}
}

func TestAddBook(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/add-book", strings.NewReader(mocBookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, AddBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// assert.Equal(t, rec.Body.String(), mocBookJSON+"\n")
		assert.Equal(t, rec.Body.String(), `{"code":"0000","message":"Success"}`+"\n")
	}
}

func TestSearchBook(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/search-book", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, SearchBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, rec.Body.String(), `{"method":"POST","page":"","where":"searchbook"}`+"\n")
	}
}
