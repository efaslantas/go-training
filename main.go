// EFA

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func OnPage(link string) string {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func main() {
	result := OnPage("https://www.bbc.co.uk/search?q=arda&d=SEARCH_PS")
	searchClassStart := "<p class=\"ssrcss-6arcww-PromoHeadline e1f5wbog6\"><span aria-hidden=\"false\">" //94000
	firstTitleIndex := strings.Index(result, searchClassStart)
	fmt.Println(firstTitleIndex)
	fmt.Println(len(searchClassStart))
	firstTitleIndex = firstTitleIndex + len(searchClassStart)
	last := len(result)
	result = result[firstTitleIndex:last]

	searchClassEnd := "</span>" //1000
	endTitleIndex := strings.Index(result, searchClassEnd)

	//firstTitleIndex := strings.Index(result, searchClass)
	fmt.Println(firstTitleIndex)
	fmt.Println(endTitleIndex)
	fmt.Println(result[0:endTitleIndex])
}
