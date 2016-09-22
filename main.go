package main

import (
	"flag"
	"fmt"

	"github.com/lukashambsch/news-parser/parser"
	"github.com/lukashambsch/news-parser/scraper"
	"github.com/lukashambsch/news-parser/store"
	"github.com/lukashambsch/news-parser/utils"
)

const defaultURL = "http://feed.omgili.com/5Rh5AMTrc4Pv/mainstream/posts/"

func main() {
	defer store.Store.Close()

	URL := flag.String("url", defaultURL, "URL to scrape")
	flag.Parse()

	err := scraper.Scrape(*URL)
	if err == nil {
		parser.ParseAllXML(utils.XMLDir)
	} else {
		fmt.Print(err.Error())
	}
	utils.Clean()
}
