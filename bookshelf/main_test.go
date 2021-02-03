package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/appleboy/gofight/v2"
)

// 시작 + health 체크
func TestSetRouter(t *testing.T) {
	r := gofight.New()

	isTest = true

	r.GET("/").
		// SetDebug(true).
		Run(serverSetup(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			// assert.Equal(t, `{"message":"Not Found"}`+"\n", r.Body.String())
			// assert.Equal(t, http.StatusNotFound, r.Code)
			assert.Equal(t, `{"message":"Unauthorized"}`+"\n", r.Body.String())
			assert.Equal(t, http.StatusUnauthorized, r.Code)
		})
}
