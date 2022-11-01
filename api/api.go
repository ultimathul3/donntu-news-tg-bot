package api

import (
	"donntu-news-tg-bot/types"
	"encoding/json"
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
		return nil, fmt.Errorf("response error (api): %s", err.Error())
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("response body read error (api): %s", err.Error())
	}

	var response *types.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("response json error (api): %s", err.Error())
	}

	return response, nil
}
