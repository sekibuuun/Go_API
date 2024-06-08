package repositories_test

import (
	"testing"

	"github.com/sekibuuun/go_api/models"
	"github.com/sekibuuun/go_api/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectCommentList(t *testing.T) {
	articleID := 1

	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range got {
		if comment.ArticleID != articleID {
			t.Errorf("ArticleID: get %d but want %d\n", comment.ArticleID, articleID)
		}
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "test message",
	}

	expectedCommentID := 3
	newComment, err := repositories.InsertComment(testDB, comment)

	if err != nil {
		t.Fatal(err)
	}
	if newComment.CommentID != expectedCommentID {
		t.Errorf("new comment id is expected %d but got %d\n", expectedCommentID, newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `delete from comments where message = ?;`

		testDB.Exec(sqlStr, comment.Message)
	})
}
