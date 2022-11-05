package api

import (
	"donntu-news-tg-bot/types"
	"encoding/json"
	"fmt"
	"net/url"
)

func SendMessage(chatId int64, text string) (*types.Response, error) {
	return request(fmt.Sprintf(`sendMessage?chat_id=%d&text=%s&parse_mode=HTML`,
		chatId, url.QueryEscape(text),
	))
}

func SendPhotoGroup(chatId int64, photoUrls []string) error {
	var inputMediaPhoto []types.InputMedia

	// no more than 10 images for 1 message (telegram api limit)
	if len(photoUrls) > 10 {
		photoUrls = photoUrls[0:10]
	}

	for _, url := range photoUrls {
		inputMediaPhoto = append(inputMediaPhoto, types.InputMedia{
			Type:  "photo",
			Media: url,
		})
	}

	json, err := json.Marshal(inputMediaPhoto)
	if err != nil {
		return fmt.Errorf("json marshal error (SendPhotoMedia): %s", err.Error())
	}

	return requestWithoutResponse(fmt.Sprintf(`sendMediaGroup?chat_id=%d&media=%s`,
		chatId, url.QueryEscape(string(json)),
	))
}
