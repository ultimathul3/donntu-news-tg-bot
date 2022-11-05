package parser

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func ParseDonntuNews(donntuNewsLink string) (string, []string, error) {
	doc, err := getDocument(donntuNewsLink)
	if err != nil {
		return "", nil, err
	}

	title := doc.Find(".page-title").First().Text()
	datetime := doc.Find("div.field:nth-child(1)").First().Children().Children().Text()
	body, err := doc.Find(".field--name-body").Children().Children().Html()
	if err != nil {
		return "", nil, fmt.Errorf("parse error: %s", err.Error())
	}

	body = formatHtml(body)

	var titleImg string
	var otherImgs []string
	doc.Find(".l-content").Find("img").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		if i == 0 {
			titleImg = fmt.Sprintf(`<a href="%s">Фото</a>`, src)
		} else {
			otherImgs = append(otherImgs, src)
		}
	})

	news := fmt.Sprintf("<strong>%s\n(%s)</strong>\n%s\n\n%s", title, datetime, titleImg, body)

	return news, otherImgs, nil
}

func ParseLastNews() (link string, datetime time.Time, err error) {
	doc, err := getDocument("http://donntu.ru/news/")
	if err != nil {
		return "", time.Time{}, err
	}

	link, _ = doc.Find(".news-main").Children().Children().Children().Attr("href")
	link = "http://donntu.ru" + link

	doc, err = getDocument(link)
	if err != nil {
		return "", time.Time{}, err
	}

	datetimeText := doc.Find("div.field:nth-child(1)").First().Children().Children().Text()
	datetime, err = time.Parse("02.01.2006 - 15:04", datetimeText)
	if err != nil {
		return "", time.Time{}, err
	}

	return link, datetime, nil
}
