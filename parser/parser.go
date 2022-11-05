package parser

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	tableRegex = `(?U)<table.*\/table>`
)

func getDocument(link string) (*goquery.Document, error) {
	response, err := http.Get(link)
	if err != nil {
		return nil, fmt.Errorf("response error (parser): %s", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error (parser): %d %s", response.StatusCode, response.Status)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, fmt.Errorf("parse error: %s", err.Error())
	}

	return doc, nil
}

// formatHtml removes tags that are not supported by the telegram api
func formatHtml(html string) string {
	html = replaceTag(html, "img", "", "")
	html = replaceTag(html, "p", "", "\n\n")
	html = replaceTag(html, "span", "", "")
	html = replaceTag(html, "br", "\n", "\n")
	html = replaceTag(html, "hr", "", "")
	html = replaceTag(html, "div", "", "\n")
	html = replaceTag(html, "noindex", "", "\n")
	html = replaceTag(html, "ol", "", "")
	html = replaceTag(html, "ul", "", "")
	html = replaceTag(html, "li", "", "\n")

	re := regexp.MustCompile(tableRegex)
	html = re.ReplaceAllString(html, "")

	for i := 0; i < 6; i++ {
		html = replaceTag(html, fmt.Sprintf("h%d", i+1), "", "\n")
	}

	return html
}

func replaceTag(html, tag, openTagReplace, closeTagReplace string) string {
	re := regexp.MustCompile(fmt.Sprintf(`(?U)<%s.*>`, tag))
	html = re.ReplaceAllString(html, openTagReplace)
	html = strings.ReplaceAll(html, fmt.Sprintf(`</%s>`, tag), closeTagReplace)

	return html
}
