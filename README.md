Parser for news posts
===========================
Parser to download, extract, parse, and store xml files found at http://feed.omgili.com/5Rh5AMTrc4Pv/mainstream/posts/.

To Run (from root project directory):

    git clone https://github.com/lukashambsch/news-parser.git
    go get github.com/PuerkitoBio/goquery
    go get github.com/garyburd/redigo/redis
    go install .
    news-parser

To Run All Tests (from root project directory):

    go test ./...
