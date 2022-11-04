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

func get(UrlWithParams string) ([]byte, error) {
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
		return body, fmt.Errorf("response body read error (api): %s", err.Error())
	}

	return body, err
}

func request(UrlWithParams string) (*types.Response, error) {
	body, err := get(UrlWithParams)
	if err != nil {
		return nil, err
	}

	var response *types.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("response json error (api): %s", err.Error())
	}

	return response, nil
}

func requestWithoutResponse(UrlWithParams string) error {
	_, err := get(UrlWithParams)
	if err != nil {
		return err
	}

	return nil
}
