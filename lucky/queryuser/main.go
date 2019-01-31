package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	//url := "https://github.com/devfeel/dotweb/issues/181"
	url := "https://github.com/mafanr/meq/issues/100"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	fmt.Println(doc, err)
	if err == nil {
		doc.Find(".timeline-comment-wrapper").Each(func(i int, s *goquery.Selection) {
			nickName := strings.TrimSpace(s.Find(".timeline-comment-header-text .css-truncate").Text())
			msg := strings.Replace(strings.Replace(s.Find("task-lists").Find("p").Text(), "\r", " ", -1), "\n", " ", -1)
			fmt.Println(nickName, "祝：", msg)
		})
	}

	fmt.Println("query html finished")
}
