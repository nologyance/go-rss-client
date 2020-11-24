package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

func main() {
	flag.Parse()
	checkDateStr := flag.Arg(0)
	layout := "2006/01/02"

	fp := gofeed.NewParser()
	urls := []string{"url1", "url2"}

	for _, url := range urls {
		feed, _ := fp.ParseURL(url)
		for _, item := range feed.Items {
			checkDate, _ := time.Parse(layout, checkDateStr)
			publishDate, _ := time.Parse(layout, item.Published)
			if checkDate.After(publishDate) {
				fmt.Println(item.Title)
				fmt.Println(item.Link)
				fmt.Println(item.Published)
				fmt.Println()

			}
		}
	}

}
