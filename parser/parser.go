package parser

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"sync"

	"github.com/lukashambsch/news-parser/models"
	"github.com/lukashambsch/news-parser/store"
)

func ingest(path string, wg *sync.WaitGroup, throttle chan int) {
	defer wg.Done()

	post := ParseXML(path)
	store.SetPost(post)
	<-throttle
}

func ParseAllXML(src string) {
	var wg sync.WaitGroup
	throttle := make(chan int, 100)

	filepath.Walk(src, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			throttle <- 1
			wg.Add(1)
			go ingest(path, &wg, throttle)
			wg.Wait()
		}
		return nil
	})
}

func ParseXML(path string) models.Post {
	var post models.Post
	f, _ := os.Open(path)
	decoder := xml.NewDecoder(f)

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "document" {
				decoder.DecodeElement(&post, &se)
			}
		}
	}

	return post
}
