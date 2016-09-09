package store

import (
	"fmt"

	"github.com/garyburd/redigo/redis"

	"github.com/lukashambsch/news-parser/models"
)

func SetPost(post models.Post) {
	_, err := Store.Do("SETNX", post.PostID, post)
	if err != nil {
		fmt.Println(err)
	}
	fromRedis, err := redis.String(Store.Do("GET", fmt.Sprintf("post:%s", post.PostID)))
	if err != nil {
		fmt.Println("Key not found")
	}
	fmt.Println(fromRedis)
}
