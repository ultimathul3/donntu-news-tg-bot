package handler

import (
	"donntu-news-tg-bot/api"
	"donntu-news-tg-bot/logger"
	"donntu-news-tg-bot/parser"
	"donntu-news-tg-bot/types"
	"fmt"
	"regexp"
	"strings"
)

const (
	donntuNewsLinkRegexp = `(http:\/\/)?(www\.)?donntu\.ru\/news\/id\d{12}`
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
	text = strings.ReplaceAll(text, " ", "")

	if regexp.MustCompile(donntuNewsLinkRegexp).MatchString(text) {
		if err := parser.ParseDonntuNews(text); err != nil {
			logger.Log.Info("parse error:", err.Error())
		}
	}

	response, err := api.SendMessage(chatId, text)
	if err != nil {
		logger.Log.Info("response error (updates):", err.Error())
	}
	logger.Log.Info("send message:", fmt.Sprintf("%+v", response))
}
