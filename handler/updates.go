package handler

import (
	"donntu-news-tg-bot/api"
	"donntu-news-tg-bot/logger"
	"donntu-news-tg-bot/parser"
	"donntu-news-tg-bot/types"
	"fmt"
	"math"
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
	text = strings.ReplaceAll(text, " ", "")

	if regexp.MustCompile(donntuNewsLinkRegex).MatchString(text) {
		sendNews(chatId, text)
	}
}

func sendNews(chatId int64, url string) {
	news, images, err := parser.ParseDonntuNews(url)
	if err != nil {
		logger.Log.Info(err.Error())
		return
	}

	response, err := api.SendMessage(chatId, news)
	if err != nil {
		logger.Log.Info(err.Error())
		return
	}
	logger.Log.Info(fmt.Sprintf("send message: %+v", response))

	if len(images) > 0 && len(images) <= 20 {
		requests := 1
		stop := len(images)
		if int(math.Ceil(float64(len(images))/10.0)) > 1 {
			requests = 2
			stop = 10
		}
		for i := 0; i < requests; i++ {
			if requests == 2 && i == 1 {
				stop = len(images)
			}

			err := api.SendPhotoGroup(chatId, images[i*10:stop])
			if err != nil {
				logger.Log.Info(err.Error())
				return
			}
		}
		logger.Log.Info(fmt.Sprintf("send %d images (chat_id: %d): %s", len(images), chatId, images))
	}
}
