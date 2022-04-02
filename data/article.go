package data

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Articles struct {
	Post []Article
}

type Article struct {
	Id        uint16 `json:"id"`
	Title     string `json:"title"`
	Full_text string `json:"full_text"`
}

func connect() *sql.DB {
	db, err := sql.Open("mysql", "asaweb:asaweb@/golang")
	if err != nil {
		panic(err)
	}
	return db
}

func Insert(title string, full_text string) error {
	db := connect()
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO articles(title, full_text) VALUES(?,?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(title, full_text)

	if err != nil {
		return err
	}

	return nil
}

func SelectAll() Articles {
	db := connect()
	defer db.Close()

	query, err := db.Query("SELECT * FROM articles")
	if err != nil {
		panic(err)
	}
	defer query.Close()

	var articles Articles
	for query.Next() {
		var article Article
		query.Scan(&article.Id, &article.Title, &article.Full_text)
		articles.Post = append(articles.Post, article)
	}
	return articles
}
func FindByID(id uint16) (Article, error) {
	db := connect()
	defer db.Close()
	statement, err := db.Prepare("SELECT * FROM articles WHERE id=?")
	if err != nil {
		return Article{}, err
	}
	defer statement.Close()
	res := statement.QueryRow(id)
	if err != nil {
		return Article{}, err
	}
	var article Article
	res.Scan(&article.Id, &article.Title, &article.Full_text)

	return article, nil
}

func DeleteById(id uint16) error {
	db := connect()
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM articles WHERE id=?")
	if err != nil {
		return err
	}
	defer statement.Close()
	
	statement.Exec(id)
	return nil
}
