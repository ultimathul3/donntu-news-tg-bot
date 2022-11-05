package observer

import (
	"donntu-news-tg-bot/db"
	"donntu-news-tg-bot/handler"
	"donntu-news-tg-bot/logger"
	"donntu-news-tg-bot/parser"
	"fmt"
	"os"
	"strconv"
	"time"
)

func Observe() {
	for {
		logger.Log.Info("observer: looking for news...")

		link, datetime, err := parser.ParseLastNews()
		if err != nil {
			logger.Log.Info(err.Error())
			continue
		}

		lastNewsDatetime, err := db.GetLastNewsDatetime()
		if err != nil {
			logger.Log.Info(err.Error())
			continue
		}

		if datetime.After(lastNewsDatetime) {
			subscribers, err := db.GetAllSubscribers()
			if err != nil {
				logger.Log.Info(err.Error())
				continue
			}

			logger.Log.Info(fmt.Sprintf("observer: send news to %d subscribers...", len(subscribers)))

			news, images, err := parser.ParseDonntuNews(link)
			if err != nil {
				logger.Log.Info(err.Error())
				return
			}

			for _, subscriber := range subscribers {
				handler.SendNews(news, images, subscriber)
			}

			db.UpdateLastNews(datetime)
		} else {
			logger.Log.Info("observer: news not found")
		}

		period, err := strconv.Atoi(os.Getenv("CHECK_PERIOD"))
		if err != nil {
			logger.Log.Info(fmt.Sprintf("CHECK_PERIOD: %s", err.Error()))
			continue
		}

		time.Sleep(time.Duration(period*60) * time.Second)
	}
}
