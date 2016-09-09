package store_test

import (
	"encoding/json"
	"testing"

	"github.com/garyburd/redigo/redis"

	"github.com/lukashambsch/news-parser/models"
	"github.com/lukashambsch/news-parser/store"
)

func TestSetPost(t *testing.T) {
	store.Store.Do("FLUSHDB")

	post := models.Post{
		Type:            "type",
		Forum:           "forum",
		ForumTitle:      "forum title",
		DiscussionTitle: "discussion title",
		Language:        "language",
		GMTOffset:       "gmt offset",
		TopicURL:        "http://topicurl.com",
		TopicText:       "topic text",
		SpamScore:       .34,
		PostNum:         12,
		PostID:          "post-12",
		PostURL:         "http://posturl.com",
		PostDate:        20160101,
		PostTime:        1400,
		Username:        "username",
		Post:            "post",
		Signature:       "signature",
		ExternalLinks:   "external links",
		Country:         "country",
		MainImage:       "main image",
	}
	postJSON, _ := json.Marshal(post)
	expected := string(postJSON)

	store.SetPost(post)
	saved, _ := redis.String(store.Store.Do("LINDEX", store.NewsXMLKey, 0))

	if saved != expected {
		t.Error("Expected ", saved, " to equal ", expected)
	}
}
