package main

import (
	"os"

	"github.com/lukashambsch/news-parser/parser"
	"github.com/lukashambsch/news-parser/scraper"
	"github.com/lukashambsch/news-parser/store"
)

const URL = "http://feed.omgili.com/5Rh5AMTrc4Pv/mainstream/posts/"

func main() {
	defer store.Store.Close()
	os.RemoveAll(scraper.XMLDir)
	os.RemoveAll(scraper.ZipDir)

	//scraper.Scrape(URL)
	post := parser.ParseXML("parser/test.xml")
	store.SetPost(post)
}
