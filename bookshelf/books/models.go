package books

// Book : Book object
type Book struct {
	Name      string `json:"name" example:"Batman"`
	Author    string `json:"author" example:"Bob kane"`
	Edition   int    `json:"edition" example:"999"`
	Price     int    `json:"price" example:"14000"` // 1200 = $12.00
	Unit      string `json:"unit" example:"KRW"`    // USD, KRW
	Publisher string `json:"publisher" example:"시공사"`
}

// MockBook : Book data mock up
type MockBook struct {
	Method string `json:"method" example:"GET"`
	Page   string `json:"page" example:"1"`
	Where  string `json:"where" example:"getbook"`
}
