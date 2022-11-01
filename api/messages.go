package api

import (
	"donntu-news-tg-bot/types"
	"fmt"
)

func SendMessage(chatId int64, text string) (*types.Response, error) {
	return request(fmt.Sprintf("sendMessage?chat_id=%d&text=123%s&parse_mode=HTML",
		chatId, text,
	))
}
