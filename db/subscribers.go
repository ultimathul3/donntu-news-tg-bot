package db

func GetAllSubscribers() ([]int64, error) {
	rows, err := db.Query("SELECT chat_id FROM subscribers WHERE is_subscribed=true;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscribers []int64
	for rows.Next() {
		var subscriber int64

		err := rows.Scan(&subscriber)
		if err != nil {
			return nil, err
		}

		subscribers = append(subscribers, subscriber)
	}

	return subscribers, nil
}

func ChangeSubscribe(chatId int64, isSubscribed bool) error {
	rows, err := db.Query("SELECT chat_id FROM subscribers WHERE chat_id=$1;", chatId)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		var chatId int64
		err := rows.Scan(&chatId)
		if err != nil {
			return err
		}
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
