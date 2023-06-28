package core

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strings"
)

type Summarize struct {
	Url string `json:"url"`
}

func isTargetTag(tag *goquery.Selection) bool {
	targetTags := []string{"p", "code", "h1", "h2", "h3", "h4", "h5"}
	tagName := strings.ToLower(tag.Get(0).Data)
	for _, targetTag := range targetTags {
		if tagName == targetTag {
			return true
		}
	}
	return false
}

func (s *Summarize) Run() (string, error) {
	response, err := http.Get(s.Url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	html, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(html)))
	if err != nil {
		return "", err
	}

	var resText strings.Builder
	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		if isTargetTag(s) {
			resText.WriteString(s.Text())
			resText.WriteString("\n")
		}
	})

	return resText.String(), nil
}
