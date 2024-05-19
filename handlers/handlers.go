package handlers

import (
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List\n")
}

func ArticleNumberHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article No.1\n")
}

func ArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice\n")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
