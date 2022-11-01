package parser

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ParseDonntuNews(donntuNewsLink string) error {
	response, err := http.Get(donntuNewsLink)
	if err != nil {
		return fmt.Errorf("response error (parser): %s", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("status code error (parser): %d %s", response.StatusCode, response.Status)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return fmt.Errorf("parse error: %s", err.Error())
	}

	title := doc.Find(".page-title").First().Text()
	datetime := doc.Find("div.field:nth-child(1)").First().Children().Children().Text()

	fmt.Println(title)
	fmt.Println(datetime)

	return nil
}
