package main

import (
	"donntu-news-tg-bot/types"
	"encoding/json"
	"fmt"
	"io"
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
}
