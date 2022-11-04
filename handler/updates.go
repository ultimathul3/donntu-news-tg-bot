package handler

import (
	"donntu-news-tg-bot/api"
	"donntu-news-tg-bot/logger"
	"donntu-news-tg-bot/types"
	"fmt"
	"regexp"
	"strings"
)

const (
	donntuNewsLinkRegex = `(http:\/\/)?(www\.)?donntu\.ru\/news\/id\d{12}`
)

func handleUpdate(update types.Update) {
	var text string
	var chatId int64

	if update.Message != nil {
		text = update.Message.Text
		chatId = update.Message.Chat.Id
	} else if update.Channel_post != nil {
		text = update.Channel_post.Text
		chatId = update.Channel_post.Chat.Id
	} else {
		return
	}

	text = strings.Split(text, "@")[0]
	command := strings.Split(text, " ")

	if command[0] == "/новость" {
		if len(command) < 2 {
			sendInvalidLinkMessage(chatId)
			return
		}

		url := command[1]

		if regexp.MustCompile(donntuNewsLinkRegex).MatchString(url) {
			sendNews(chatId, url)
		} else {
			sendInvalidLinkMessage(chatId)
		}
	} else if command[1] == "/подписаться" {
		return
	} else if command[1] == "/отписаться" {
		return
	}
}

func sendInvalidLinkMessage(chatId int64) {
	response, err := api.SendMessage(chatId, "Некорректная ссылка\nНеобходимый формат: http://donntu.ru/news/idXXXXXXXXXXXX")
	if err != nil {
		logger.Log.Info(err.Error())
		return
	}
	logger.Log.Info(fmt.Sprintf("send message: %+v", response))
}
