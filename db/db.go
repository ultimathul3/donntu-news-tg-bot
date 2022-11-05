package db

import (
	"database/sql"

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

	return nil
}

func ChangeSubscribe(chatId int64, isSubscribed bool) error {
	rows, err := db.Query("SELECT * FROM subscribers WHERE chat_id=$1;", chatId)
	if err != nil {
		return err
	}

	if rows.Next() {
		var chatId int64

		rows.Scan(&chatId, nil)
		rows.Close()

		_, err = db.Exec("UPDATE subscribers SET is_subscribed=$1 WHERE chat_id=$2", isSubscribed, chatId)
		if err != nil {
			return err
		}
	} else {
		_, err = db.Exec("INSERT INTO subscribers VALUES ($1, $2);", chatId, isSubscribed)
		if err != nil {
			return err
		}
	}

	return nil
}
