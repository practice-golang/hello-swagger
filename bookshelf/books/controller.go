package books // import "books"

import (
	"net/http"

	"commons"

	"github.com/labstack/echo/v4"
)

// GetBook : Get book data
// @Tags Bookshelf
// @Security ApiKeyAuth
// @Summary Get book Data
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} MockBook "Sucess"
// @Failure 400 {object} commons.ErrorMSG "Bad Requests"
// @Router /get-book/{id} [get]
func GetBook(c echo.Context) error {
	page := c.Param("page")

	result := MockBook{
		Method: c.Request().Method,
		Page:   page,
		Where:  "getbook",
	}

	return c.JSON(http.StatusOK, result)
}

// AddBook : Add a book
// @Tags Bookshelf
// @Security ApiKeyAuth
// @Summary Add a new book
// @Accept json
// @Produce json
// @Param BookInfo body Book true "Book Data"
// @Success 200 {object} commons.ResultMSG "Sucess"
// @Failure 400 {object} commons.ErrorMSG "Bad Requests"
// @Router /add-book [post]
func AddBook(c echo.Context) error {
	// u := new(commons.Book)
	// if err := c.Bind(u); err != nil {
	// 	return err
	// }

	result := commons.ResultMSG{
		Code:    "0000",
		Message: "Success",
	}

	return c.JSON(http.StatusOK, result)
}

// SearchBook : Search book
// @Tags Bookshelf
// @Security ApiKeyAuth
// @Summary Search book
// @Accept json
// @Produce json
// @Param BookInfo body Book true "Book Data"
// @Success 200 {string} MockBook "Success"
// @Failure 400 {object} commons.ErrorMSG "Bad Requests"
// @Router /search-book [post]
func SearchBook(c echo.Context) error {
	result := MockBook{
		Method: c.Request().Method,
		Page:   "",
		Where:  "searchbook",
	}

	return c.JSON(http.StatusOK, result)
}
