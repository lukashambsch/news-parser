package store

type Document struct {
	Type            string  `xml:"type"`
	Forum           string  `xml:"forum"`
	ForumTitle      string  `xml:"forum_title"`
	DiscussionTitle string  `xml:"discussion_title"`
	Language        string  `xml:"language"`
	GMTOffset       string  `xml:"gmt_offset"`
	TopicURL        string  `xml:"topic_url"`
	TopicText       string  `xml:"topic_text"`
	SpamScore       float32 `xml:"spam_score"`
	PostNum         int     `xml:"post_num"`
	PostID          string  `xml:"post_id"`
	PostURL         string  `xml:"post_url"`
	PostDate        int     `xml:"post_date"`
	PostTime        int     `xml:"post_time"`
	Username        string  `xml:"username"`
	Post            string  `xml:"post"`
	Signature       string  `xml:"signature"`
	ExternalLinks   string  `xml:"external_links"`
	Country         string  `xml:"country"`
	MainImage       string  `xml:"main_image"`
}
