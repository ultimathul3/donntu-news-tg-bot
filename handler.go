package main

import (
	"donntu-news-tg-bot/api"
	"donntu-news-tg-bot/types"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Telegram-Bot-Api-Secret-Token") != os.Getenv("SECRET_TOKEN") {
		fileLog.Info("invalid secret token from request")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fileLog.Info("body read error:", err)
		return
	}

	var update types.Update
	err = json.Unmarshal(body, &update)
	if err != nil {
		fileLog.Info("json error:", err)
		return
	}

	fileLog.Info("update:", fmt.Sprintf("%+v", update))

	updateHandler(update)
}

func updateHandler(update types.Update) {
	var text string
	var chatId int64
	if update.Message != nil {
		text = update.Message.Text
		chatId = update.Message.Chat.Id
	} else if update.Channel_post != nil {
		text = update.Channel_post.Text
		chatId = update.Channel_post.Chat.Id
	}

	response, err := api.SendMessage(chatId, text)
	if err != nil {
		log.Fatal(err)
	}
	fileLog.Info("send message:", fmt.Sprintf("%+v", response))
}
