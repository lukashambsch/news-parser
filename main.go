package main

import (
	//"os"

	"github.com/lukashambsch/news-parser/parser"
	"github.com/lukashambsch/news-parser/scraper"
)

const URL = "http://feed.omgili.com/5Rh5AMTrc4Pv/mainstream/posts/"

func main() {
	//os.RemoveAll(scraper.XMLDir)
	//os.RemoveAll(scraper.ZipDir)

	//scraper.Scrape(URL)
	parser.ParseXML(scraper.XMLDir + "/fff067a03ddd6953c090621c4fa137a0.xml")
}
