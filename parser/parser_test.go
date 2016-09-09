package parser_test

import (
	"strings"
	"testing"

	"github.com/lukashambsch/news-parser/models"
	"github.com/lukashambsch/news-parser/parser"
)

var expected models.Post = models.Post{
	Type:            "mainstream",
	Forum:           "http://www.dailymail.co.uk/sport/index.rss",
	ForumTitle:      "Sport | Mail Online",
	DiscussionTitle: "Inside Wing's: the Chinese restaurant that's a favourite haunt for Premier League stars",
	Language:        "english",
	GMTOffset:       "",
	TopicURL:        "http://www.dailymail.co.uk/sport/football/article-3774270/Inside-Wing-s-Chinese-restaurant-s-favourite-haunt-Premier-League-stars.html",
	TopicText:       "Former United boss Louis van Gaal swore by it, Mario Balotelli used to get food delivered to his Liverpool home from it",
	SpamScore:       0,
	PostNum:         39,
	PostID:          "post-39",
	PostURL:         "http://www.dailymail.co.uk/sport/football/article-3774270/Inside-Wing-s-Chinese-restaurant-s-favourite-haunt-Premier-League-stars.html",
	PostDate:        20160905,
	PostTime:        1451,
	Username:        "Weller-Smith",
	Post:            "I hate international breaks because we have to face headline stories like this crock of shaving cream",
	Signature:       "\n\n",
	ExternalLinks:   "\n\n",
	Country:         "GB",
	MainImage:       "http://i.dailymail.co.uk/i/pix/2016/09/05/13/37EE5E8800000578-0-image-a-58_1473078496374.jpg",
}

func TestParserType(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if doc.Type != expected.Type {
		t.Error("Expected ", doc.Type, " to equal ", expected.Type)
	}
}

func TestParserForum(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if doc.Forum != expected.Forum {
		t.Error("Expected ", doc.Forum, " to equal ", expected.Forum)
	}
}

func TestParserForumTitle(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if doc.ForumTitle != expected.ForumTitle {
		t.Error("Expected ", doc.ForumTitle, " to equal ", expected.ForumTitle)
	}
}

func TestParserDiscussionTitle(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if doc.DiscussionTitle != expected.DiscussionTitle {
		t.Error("Expected ", doc.DiscussionTitle, " to equal ", expected.DiscussionTitle)
	}
}

func TestParserLanguage(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if doc.Language != expected.Language {
		t.Error("Expected ", doc.Language, " to equal ", expected.Language)
	}
}

func TestParserGMTOffset(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if doc.GMTOffset != expected.GMTOffset {
		t.Error("Expected ", doc.GMTOffset, " to equal ", expected.GMTOffset)
	}
}

func TestParserTopicURL(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if doc.TopicURL != expected.TopicURL {
		t.Error("Expected ", doc.TopicURL, " to equal ", expected.TopicURL)
	}
}

func TestParserTopicText(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if !strings.Contains(doc.TopicText, expected.TopicText) {
		t.Error("Expected ", doc.TopicText, " to contain ", expected.TopicText)
	}
}

func TestParserSpamScore(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if doc.SpamScore != expected.SpamScore {
		t.Error("Expected ", doc.SpamScore, " to equal ", expected.SpamScore)
	}
}

func TestParserPostNum(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if doc.PostNum != expected.PostNum {
		t.Error("Expected ", doc.PostNum, " to equal ", expected.PostNum)
	}
}

func TestParserPostID(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if doc.PostID != expected.PostID {
		t.Error("Expected ", doc.PostID, " to equal ", expected.PostID)
	}
}

func TestParserPostURL(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if doc.PostURL != expected.PostURL {
		t.Error("Expected ", doc.PostURL, " to equal ", expected.PostURL)
	}
}

func TestParserPostDate(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if doc.PostDate != expected.PostDate {
		t.Error("Expected ", doc.PostDate, " to equal ", expected.PostDate)
	}
}

func TestParserPostTime(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if doc.PostTime != expected.PostTime {
		t.Error("Expected ", doc.PostTime, " to equal ", expected.PostTime)
	}
}

func TestParserUsername(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if doc.Username != expected.Username {
		t.Error("Expected ", doc.Username, " to equal ", expected.Username)
	}
}

func TestParserPost(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if !strings.Contains(doc.Post, expected.Post) {
		t.Error("Expected ", doc.Post, " to contain ", expected.Post)
	}
}

func TestParserSignature(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if doc.Signature != expected.Signature {
		t.Error("Expected ", doc.Signature, " to equal ", expected.Signature)
	}
}

func TestParserExternalLinks(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if doc.ExternalLinks != expected.ExternalLinks {
		t.Error("Expected ", doc.ExternalLinks, " to equal ", expected.ExternalLinks)
	}
}

func TestParserCountry(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if doc.Country != expected.Country {
		t.Error("Expected ", doc.Country, " to equal ", expected.Country)
	}
}

func TestParserMainImage(t *testing.T) {
	doc := parser.ParseXML("test.xml")

	if doc.MainImage != expected.MainImage {
		t.Error("Expected ", doc.MainImage, " to equal ", expected.MainImage)
	}
}
