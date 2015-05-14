package model
import (
	"gopkg.in/validator.v2"
	"log"
	"time"
)

type Article struct {
	Id int `json:"id"`
	Title string `json:"title" validate:"min=4,max=64,nonzero"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func NewArticle(title string, content string) Article {
	return Article {
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}
}

func ValidArticle(article *Article) bool {
	err := validator.Validate(article)
	if err != nil {
		errs, _ := err.(validator.ErrorMap)
		for k, v := range errs {
			log.Printf("field %s is invalid from %v", k, v)
		}
		return false
	}
	return true
}