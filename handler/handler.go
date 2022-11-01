package handler

import (
	"donntu-news-tg-bot/logger"
	"donntu-news-tg-bot/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Telegram-Bot-Api-Secret-Token") != os.Getenv("SECRET_TOKEN") {
		logger.Log.Info(fmt.Sprintf("invalid secret token from request (%s)", r.RemoteAddr))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log.Info("body read error (handler):", err)
		return
	}

	var update types.Update
	err = json.Unmarshal(body, &update)
	if err != nil {
		logger.Log.Info("json error (handler):", err)
		return
	}

	logger.Log.Info("update:", fmt.Sprintf("%+v", update))

	handleUpdate(update)
}
