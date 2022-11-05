package db

import (
	"errors"
	"time"
)

func UpdateLastNews(datetime time.Time) error {
	_, err := db.Exec("UPDATE last_news SET datetime=$1;", datetime)
	if err != nil {
		return err
	}

	return nil
}

func GetLastNewsDatetime() (time.Time, error) {
	rows, err := db.Query("SELECT datetime FROM last_news;")
	if err != nil {
		return time.Time{}, err
	}
	defer rows.Close()

	var datetime time.Time
	if rows.Next() {
		err := rows.Scan(&datetime)
		if err != nil {
			return time.Time{}, err
		}
	} else {
		return time.Time{}, errors.New("db: datetime not set")
	}

	return datetime, nil
}
