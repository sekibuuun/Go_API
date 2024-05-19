package main

import (
	"log"
	"net/http"

	"github.com/sekibuuun/go_api/handlers"
)

func main() {

	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleNumberHandler)
	http.HandleFunc("/article/nice", handlers.ArticleNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
