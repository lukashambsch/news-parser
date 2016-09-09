package store

import (
	"encoding/json"
	"fmt"

	"github.com/lukashambsch/news-parser/models"
)

const newsXMLKey = "NEWS_XML"

func SetPost(post models.Post) {
	postJSON, _ := json.Marshal(post)
	Store.Do("LREM", newsXMLKey, 1, string(postJSON))
	_, err := Store.Do("LPUSH", newsXMLKey, string(postJSON))
	if err != nil {
		fmt.Println(err)
	}
}
