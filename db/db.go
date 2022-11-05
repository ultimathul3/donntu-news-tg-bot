package db

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func Init() error {
	var err error

	db, err = sql.Open("sqlite3", "database.db")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS subscribers (chat_id INTEGER, is_subscribed BOOLEAN);")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS last_news (datetime DATETIME);")
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT * FROM last_news;")
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		_, err = db.Exec("INSERT INTO last_news VALUES ($1);", time.Now())
		if err != nil {
			return err
		}
	}

	return nil
}
