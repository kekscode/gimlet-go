package gimlet

/*
Implements a basic blog API with CRUD operations
*/

import (
	"time"

	"github.com/pborman/uuid"
)

/* Article struct for storing our articles */
type Article struct {
	ID        uuid.UUID
	Title     string
	Subtitle  string
	Body      string
	Timestamp string
}

/* all artciles */
var content Article

func returnDateString() string {
	now := time.Now()
	t := now.Format(time.RFC3339)
	return t
}

func createOrUpdateArticle(title, subtitle, body string, timestamp string) Article {

	post := Article{
		Title:     title,
		Subtitle:  subtitle,
		Body:      body,
		Timestamp: timestamp,
	}
	return post
}

func fetchArticle(id uuid.UUID)  {}
func fetchAllArticles()          {}
func createArticle(id uuid.UUID) {}
func updateArticle(id uuid.UUID) {}
func deleteArticle(id uuid.UUID) {}
