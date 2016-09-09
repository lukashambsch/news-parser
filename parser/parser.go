package parser

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"strings"

	"github.com/lukashambsch/news-parser/models"
	"github.com/lukashambsch/news-parser/store"
)

func ParseAllXML(src string) {
	filepath.Walk(src, func(path string, f os.FileInfo, err error) error {
		if strings.Contains(path, ".xml") {
			post := ParseXML(path)
			store.SetPost(post)
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
