package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Comment struct {
	CommentID int       `json:"comment_id"`
	ArticleId int       `json:"article_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type Article struct {
	ID          int       `json:"article_id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	UserName    string    `json:"user_name"`
	NiceNum     int       `json:"nice"`
	CommentList []Comment `json:"comments"`
	CreatedAt   time.Time `json:"created_at"`
}

func main() {
	comment1 := Comment{
		CommentID: 1,
		ArticleId: 1,
		Message:   "test1",
		CreatedAt: time.Now(),
	}
	comment2 := Comment{
		CommentID: 2,
		ArticleId: 1,
		Message:   "test2",
		CreatedAt: time.Now(),
	}
	article := Article{
		ID:          1,
		Title:       "first title",
		Contents:    "こんにちは",
		UserName:    "aaa",
		NiceNum:     1,
		CommentList: []Comment{comment1, comment2},
		CreatedAt:   time.Now(),
	}

	jsonData, err := json.Marshal(article)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", jsonData)
}
