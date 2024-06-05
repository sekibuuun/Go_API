package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sekibuuun/go_api/models"
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `insert into articles (title, contents, username, nice, created_at) values (?, ?, ?, 0, now());`

	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)

	if err != nil {
		fmt.Println(err)
		return models.Article{}, err
	}

	id, _ := result.LastInsertId()

	newArticle.ID = int(id)

	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `select article_id, title, contents, username, nice from articles limit ? offset ?;`
	// limit 取得件数, offset 取得開始位置

	articleArray := make([]models.Article, 0)

	rows, err := db.Query(sqlStr, 5, (page-1)*5)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)

		if err != nil {
			fmt.Println(err)
			return nil, err
		} else {
			articleArray = append(articleArray, article)
		}
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `select * from articles where article_id = ?;`

	row := db.QueryRow(sqlStr, articleID)

	if err := row.Err(); err != nil {
		fmt.Println(err)
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime

	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)

	if err != nil {
		fmt.Println(err)
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err)
		return err
	}

	const sqlGetNice = `select nice from articles where article_id = ?;`

	const sqlUpdateNice = `update articles set nice = ? where article_id = ?;`

	row := tx.QueryRow(sqlGetNice, articleID)

	if err := row.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	var niceNum int
	err = row.Scan(&niceNum)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(sqlUpdateNice, niceNum+1, articleID)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
