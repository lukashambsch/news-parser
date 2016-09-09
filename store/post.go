package store

import (
	"fmt"

	"github.com/lukashambsch/news-parser/models"
)

func SetPost(post models.Post) {
	key := fmt.Sprintf("post:%s", post.PostURL)
	_, err := Store.Do("SETNX", key, post)
	if err != nil {
		fmt.Println(err)
	}
}
