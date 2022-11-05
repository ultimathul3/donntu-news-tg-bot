package handler

import (
	"donntu-news-tg-bot/api"
	"donntu-news-tg-bot/logger"
	"fmt"
	"math"
)

func SendNews(news string, images []string, chatId int64) {
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
