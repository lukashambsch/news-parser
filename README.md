Parser to download, extract, parse, and store xml files found at http://feed.omgili.com/5Rh5AMTrc4Pv/mainstream/posts/.

To Run:
git clone https://github.com/lukashambsch/news-parser.git
go get github.com/PuerkitoBio/goquery
go get github.com/garyburd/redigo/redis
# path starts from $GOPATH/src
go install /path/to/news-parser
news-parser
