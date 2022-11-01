package api

import (
	"donntu-news-tg-bot/types"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	tgApi = "https://api.telegram.org/"
)

func request(UrlWithParams string) (*types.Response, error) {
	accessToken := os.Getenv("ACCESS_TOKEN")

	r, err := http.Get(tgApi + fmt.Sprintf("bot%s/%s",
		accessToken, UrlWithParams,
	))
	if err != nil {
		return nil, errors.New("response error: " + err.Error())
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("response body read error: " + err.Error())
	}

	var response *types.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.New("response json error: " + err.Error())
	}

	return response, nil
}
