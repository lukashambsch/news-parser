package store

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

var Store redis.Conn

func init() {
	var err error

	Store, err = redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
}
