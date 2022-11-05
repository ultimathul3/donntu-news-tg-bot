package handler

import (
	"donntu-news-tg-bot/api"
	"donntu-news-tg-bot/db"
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
	} else if command[0] == "/подписаться" {
		subscribe(chatId)
	} else if command[0] == "/отписаться" {
		unsubscribe(chatId)
	}
}

func subscribe(chatId int64) {
	err := db.ChangeSubscribe(chatId, true)
	if err != nil {
		logger.Log.Info(fmt.Sprintf("db error: %s", err.Error()))
		return
	}
	logger.Log.Info(fmt.Sprintf("db: chat_id %d subscribed", chatId))

	response, err := api.SendMessage(chatId, "Вы подписались на рассылку новостей")
	if err != nil {
		logger.Log.Info(err.Error())
		return
	}
	logger.Log.Info(fmt.Sprintf("send message: %+v", response))
}

func unsubscribe(chatId int64) {
	err := db.ChangeSubscribe(chatId, false)
	if err != nil {
		logger.Log.Info(fmt.Sprintf("db error: %s", err.Error()))
		return
	}
	logger.Log.Info(fmt.Sprintf("db: chat_id %d unsubscribed", chatId))

	response, err := api.SendMessage(chatId, "Вы отписались от рассылки новостей")
	if err != nil {
		logger.Log.Info(err.Error())
		return
	}
	logger.Log.Info(fmt.Sprintf("send message: %+v", response))
}

func sendInvalidLinkMessage(chatId int64) {
	response, err := api.SendMessage(chatId, "Некорректная ссылка\nНеобходимый формат: http://donntu.ru/news/idXXXXXXXXXXXX")
	if err != nil {
		logger.Log.Info(err.Error())
		return
	}
	logger.Log.Info(fmt.Sprintf("send message: %+v", response))
}
