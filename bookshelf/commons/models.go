package commons // import "commons"

import "time"

// ResultMSG : 결과
type ResultMSG struct {
	Code    string `json:"code" example:"0000"`
	Message string `json:"message" example:"Sucess"`
}

// Error : Swagger error sample
type Error struct {
	ErrorCode    int
	ErrorMessage string
	CreatedAt    time.Time
}

// ErrorMSG : Swagger error sample
type ErrorMSG struct {
	Message string `json:"message" example:"missing or malformed jwt"`
}
